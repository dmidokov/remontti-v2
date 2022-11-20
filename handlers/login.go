package handlers

import (
	"encoding/json"
	"github.com/dmidokov/remontti-v2/companyservice"
	"github.com/dmidokov/remontti-v2/translationservice"
	"github.com/dmidokov/remontti-v2/userservice"
	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
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
	if r.Method == "POST" {
		hm.loginPOST(w, r)
	} else {
		// TODO: проверить используется ли этот кусок?
		if hm.Config.MODE == "dev" {
			if r.Method == "OPTIONS" {
				setCorsHeaders(&w, r)
				return
			}
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}

func (hm *HandlersModel) loginPOST(w http.ResponseWriter, r *http.Request) {

	var log = hm.Logger

	if hm.Config.MODE == "dev" {
		setCorsHeaders(&w, r)
	}

	host := r.Host

	var companies = companyservice.CompanyModel{DB: hm.DB}
	var translation = translationservice.TranslationsModel{DB: hm.DB}
	var users = userservice.UserModel{DB: hm.DB}

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
		log.Warning("Не удалось декодировать JSON: %s", err)
		err := json.NewEncoder(w).Encode(response{
			Status:  "error",
			Message: pageData.Translation["InvalidData"],
			Errors:  []string{"Invalid data"}})

		if err != nil {
			log.Error("Не удалось кодировать JSON: %s", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}

		return
	}

	if len(v.Login) > 0 && len(v.Password) > 0 {

		// НЕПРАВИЛЬНОЕ ОПИСАНИЕ ОШИБКИ ИСПРАВИТЬ
		company, err := companies.GetCompanyByHostName(host)
		if err != nil {
			log.Warning("Компании не найдены: %s", err)
			err := json.NewEncoder(w).Encode(response{
				Status:  "error",
				Message: pageData.Translation["ErrorTryAgain"],
				Errors:  []string{"Internal server error"}})

			if err != nil {
				log.Error("Не удалось кодировать JSON: %s", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}

			return
		}

		user, err := users.GetByNameAndCompanyId(v.Login, company.CompanyId)

		if err != nil {
			if err == pgx.ErrNoRows {
				log.Warning("Пользователь не найден: %s", err)
				err := json.NewEncoder(w).Encode(response{
					Status:  "error",
					Message: pageData.Translation["UserIsNotExists"],
					Errors:  []string{"User is not exists"}})

				if err != nil {
					log.Error("Не удалось кодировать JSON: %s", err)
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				}

				return
			}
			log.Error("Ошибка при получении пользовательских данных: %s", err)
			err := json.NewEncoder(w).Encode(response{
				Status:  "error",
				Message: pageData.Translation["ErrorTryAgain"],
				Errors:  []string{"Internal server error"}})

			if err != nil {
				log.Error("Не удалось кодировать JSON: %s", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}

			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(v.Password))

		if err != nil {
			log.Warning("Неверный пароль: %s", err)
			err := json.NewEncoder(w).Encode(response{
				Status:  "error",
				Message: pageData.Translation["InvalidUserOrPassword"],
				Errors:  []string{"Invalid password"}})

			if err != nil {
				log.Error("Не удалось кодировать JSON: %s", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}

			return
		} else {

			session, _ := hm.CookieStore.Get(r, hm.Config.SESSIONS_SECRET)
			session.Values["authenticated"] = true
			session.Values["userid"] = user.Id
			session.Values["companyid"] = company.CompanyId
			session.Options.MaxAge = 3600
			err := session.Save(r, w)

			if err != nil {
				err := json.NewEncoder(w).Encode(response{
					Status:  "error",
					Message: pageData.Translation["ErrorTryAgain"],
					Errors:  []string{"Internal server error"}})

				if err != nil {
					log.Error("Не удалось кодировать JSON: %s", err)
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				}

				return
			}

			err = json.NewEncoder(w).Encode(response{
				Status: "ok",
				Errors: []string{},
			})

			if err != nil {
				log.Error("Не удалось кодировать JSON: %s", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
			println("login.......................")
			http.Redirect(w, r, "/", http.StatusSeeOther)
			//return
		}

	} else {
		log.Warning("Логин или пароль пустой")
		err := json.NewEncoder(w).Encode(response{
			Status:  "error",
			Message: pageData.Translation["EmptyLoginOrPassword"],
			Errors:  []string{"Login or password is empty"}})

		if err != nil {
			log.Error("Не удалось кодировать JSON: %s", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

}

func (hm *HandlersModel) logout(w http.ResponseWriter, r *http.Request) {

	var log = hm.Logger

	log.Info("Логаут")

	if r.Method != http.MethodGet {
		log.Error("Неверный метод")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	session, _ := hm.CookieStore.Get(r, hm.Config.SESSIONS_SECRET)
	session.Values["authenticated"] = false
	session.Options.MaxAge = -1
	err := session.Save(r, w)

	if err != nil {
		log.Error("Не удалось сохранить данные сессии %s", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	w.Header().Set("cache-control", "none")
	http.Redirect(w, r, "https://control.remontti.site/", http.StatusSeeOther)
}
