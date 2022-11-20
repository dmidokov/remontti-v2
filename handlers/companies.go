package handlers

import (
	"encoding/json"
	"github.com/dmidokov/remontti-v2/userservice"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/dmidokov/remontti-v2/companyservice"
)

type companiesData struct {
	ID   int
	Host string
	Name string
}

type companiesPageData struct {
	Title       string
	Translation map[string]string
	Companies   []companiesData
	Navigation  map[string]navigationData
}

type newCompanyForm struct {
	Name string `json:"name"`
	Host string `json:"host"`
}

type CompaniesResult struct {
	CompanyName string `json:"company"`
	HostName    string `json:"host"`
	CompanyId   int    `json:"id"`
}

func (hm *HandlersModel) companies(w http.ResponseWriter, r *http.Request) {

	var rootPath = hm.Config.ROOT_PATH + "/web/ui/"

	files := []string{
		rootPath + "companies.page.gohtml",
		rootPath + "base.layout.gohtml",
		rootPath + "footers/footer.partial.gohtml",
		rootPath + "headers/mainpage.partial.gohtml",
		rootPath + "bodies/companies.partial.gohtml",
		rootPath + "navigations/topnavigation.partial.gohtml",
		rootPath + "heads/companies.partial.gohtml",
	}

	var pageData = companiesPageData{
		Title:       "",
		Translation: make(map[string]string),
		Companies:   make([]companiesData, 0),
		Navigation:  map[string]navigationData{},
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	pageData.Translation, err = hm.getTranslations("mainpage", "navigation", "companies")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	sessionsData, err := hm.getSessionData(r)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	pageData.Navigation, err = hm.getUserNavigation(sessionsData.UserId, pageData.Translation)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	companies, err := hm.getCompanies()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	pageData.Companies = *companies

	err = ts.Execute(w, pageData)
	if err != nil {
		log.Print(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

func (hm *HandlersModel) addCompany(w http.ResponseWriter, r *http.Request) {

	var companyForm *newCompanyForm
	err := json.NewDecoder(r.Body).Decode(&companyForm)

	if err != nil {
		log.Printf("Invalid data: %s", err)
		json.NewEncoder(w).Encode(response{
			Status:  "error",
			Message: pageData.Translation["InvalidData"],
			Errors:  []string{"Invalid data"}})
		return
	}

	if (len(companyForm.Name) < 2) || (len(strings.Split(companyForm.Host, ".")) != 3) {
		log.Printf("Invalid data:")
		json.NewEncoder(w).Encode(response{
			Status:  "error",
			Message: pageData.Translation["InvalidData"],
			Errors:  []string{"Invalid data"}})
		return
	}

	var companyService = companyservice.CompanyModel{DB: hm.DB}
	_, err = companyService.Add(companyForm.Name, companyForm.Host)

	if err != nil {
		log.Printf("Internal server error: %s", err)
		json.NewEncoder(w).Encode(response{
			Status:  "error",
			Message: pageData.Translation["Internal server error"],
			Errors:  []string{"Internal server error"}})
		return
	}

	json.NewEncoder(w).Encode(response{
		Status: "ok",
		Errors: []string{},
	})

}

func (hm *HandlersModel) getCompaniesApi(w http.ResponseWriter, r *http.Request) {
	log := hm.Logger
	log.Info("Get companies")

	var companiesService = companyservice.CompanyModel{DB: hm.DB}
	var userService = userservice.UserModel{DB: hm.DB, CookieStore: hm.CookieStore}

	// пользователь идентифицируется по сессии, это не очень хорошо, так как
	// не позволяет сделать запрос не из браузера, то есть другие приложения не
	// смогут запросить список компаний и другие вещи при необходимости
	// вынуждены будут как-то проходить авторизацию или хз что еще, стоит
	// завести для методов принимающих запросы из вне обращаться по какому-то токену
	userId, err := userService.GetCurrentUserId(r, hm.Config.SESSIONS_SECRET)
	// надо ли возвращать тут json? или стоит сразу кидать 500, ведь по сути тут произошла ошибка
	// если оставляет так, то методы в JS надо снабдить умением выкидывать сообщения об ошибках
	// в какую-нибудь всплывашку
	if err != nil {
		// TODO: Вынести подобного рода обработчики в отдельный пакет или хотя бы файл
		err := json.NewEncoder(w).Encode(response{
			Status:  "error",
			Message: pageData.Translation["ErrorTryAgain"],
			Errors:  []string{"Internal server error"}})

		if err != nil {
			log.Error("Не удалось кодировать JSON: %s", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}

	companies, err := companiesService.GetAllForUser(userId)

	result := make([]*CompaniesResult, 0)
	for _, item := range companies {
		result = append(
			result,
			&CompaniesResult{CompanyName: item.CompanyName, HostName: item.HostName, CompanyId: item.CompanyId},
		)
	}

	json.NewEncoder(w).Encode(result)

}
