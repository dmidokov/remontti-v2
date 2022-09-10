package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"text/template"

	"github.com/dmidokov/remontti-v2/companyservice"
	"github.com/dmidokov/remontti-v2/translationservice"
	"github.com/dmidokov/remontti-v2/userservice"
	"golang.org/x/crypto/bcrypt"
)

type loginPageData struct {
	Title       string
	Translation map[string]string
	Navigation  map[string]navigationData
}

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

var pageData = loginPageData{
	Title:       "",
	Translation: make(map[string]string),
	Navigation:  make(map[string]navigationData),
}

// Страница логина
func (hm *HandlersModel) login(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		hm.loginGET(w, r)
	} else if r.Method == "POST" {
		hm.loginPOST(w, r)
	}
}

func (h *HandlersModel) loginGET(w http.ResponseWriter, r *http.Request) {

	var translation translationservice.TranslationsModel = translationservice.TranslationsModel{DB: h.DB}

	var rootPath = h.Config.ROOT_PATH + "/web/ui/"

	files := []string{
		rootPath + "login.page.gohtml",
		rootPath + "login.layout.gohtml",
		rootPath + "navigations/navigationNoAuth.partial.gohtml",
		rootPath + "bodies/login.partial.gohtml",
		rootPath + "heads/login.partial.gohtml",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	pageData.Title = "Вход"

	translations, err := translation.Get("loginpage")
	if err != nil {
		println(err.Error())
	}

	for _, translation := range translations {
		pageData.Translation[translation.Label] = translation.Ru
	}

	err = ts.Execute(w, pageData)
	if err != nil {
		log.Print(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

}

func (h *HandlersModel) loginPOST(w http.ResponseWriter, r *http.Request) {

	host := r.Host

	var companies companyservice.CompanyModel = companyservice.CompanyModel{DB: h.DB}
	var translation translationservice.TranslationsModel = translationservice.TranslationsModel{DB: h.DB}
	var users userservice.UserModel = userservice.UserModel{DB: h.DB}

	translations, err := translation.Get("loginpage")
	if err != nil {
		println(err.Error())
	}

	for _, translation := range translations {
		pageData.Translation[translation.Label] = translation.Ru
	}

	var v User
	err = json.NewDecoder(r.Body).Decode(&v)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		log.Printf("Invalid data: %s", err)
		json.NewEncoder(w).Encode(response{
			Status:  "error",
			Message: pageData.Translation["InvalidData"],
			Errors:  []string{"Invalid data"}})
		return
	}

	if len(v.Login) > 0 && len(v.Password) > 0 {

		// НЕПРАВИЛЬНОЕ ОПИСАНИЕ ОШИБКИ ИСПРАВИТЬ 
		company, err := companies.GetCompanyByHostName(host)
		if err != nil {
			json.NewEncoder(w).Encode(response{
				Status:  "error",
				Message: pageData.Translation["ErrorTryAgain"],
				Errors:  []string{"Internal server error"}})
			return
		}

		user, err := users.GetByNameAndCompanyId(v.Login, company.CompanyId)

		if err != nil {
			if err == sql.ErrNoRows {
				log.Printf("User is not exists with error: %s", err)
				json.NewEncoder(w).Encode(response{
					Status:  "error",
					Message: pageData.Translation["UserIsNotExists"],
					Errors:  []string{"User is not exists"}})
				return
			}
			log.Printf("Error scaning user password from DB response with error: %s", err)
			json.NewEncoder(w).Encode(response{
				Status:  "error",
				Message: pageData.Translation["ErrorTryAgain"],
				Errors:  []string{"Internal server error"}})
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(v.Password))

		if err != nil {
			log.Printf("Invalid password: %s", err)
			json.NewEncoder(w).Encode(response{
				Status:  "error",
				Message: pageData.Translation["InvalidUserOrPassword"],
				Errors:  []string{"Invalid password"}})
			return
		} else {

			session, _ := h.CookieStore.Get(r, "session-key")
			session.Values["authenticated"] = true
			session.Values["userid"] = user.Id
			session.Values["companyid"] = company.CompanyId
			session.Options.MaxAge = 3600
			session.Save(r, w)

			json.NewEncoder(w).Encode(response{
				Status: "ok",
				Errors: []string{},
			})

			return
		}

	} else {
		log.Printf("Login or password is empty")
		json.NewEncoder(w).Encode(response{
			Status:  "error",
			Message: pageData.Translation["EmptyLoginOrPassword"],
			Errors:  []string{"Login or password is empty"}})
		return
	}

}

func (hm *HandlersModel) logout(w http.ResponseWriter, r *http.Request) {
	session, _ := hm.CookieStore.Get(r, "session-key")
	session.Values["authenticated"] = false
	session.Options.MaxAge = -1
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
