package handlers

import (
	"encoding/json"
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
		http.Error(w, "Internal Server Error", 500)
		return
	}

	pageData.Translation, err = hm.getTranslations("mainpage", "navigation", "companies")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

	sessionsData, err := hm.getSessionData(r)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

	pageData.Navigation, err = hm.getUserNavigation(sessionsData.UserId, pageData.Translation)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

	companies, err := hm.getCompanies()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
	pageData.Companies = *companies

	err = ts.Execute(w, pageData)
	if err != nil {
		log.Print(err)
		http.Error(w, "Internal Server Error", 500)
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
