package handlers

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/dmidokov/remontti-v2/database"
	"golang.org/x/crypto/bcrypt"
)

type loginPageData struct {
	Title       string
	Translation map[string]string
}

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type response struct {
	Status  string   `json:"status"`
	Errors  []string `json:"errors"`
	Message string   `json:"message" `
}

var pageData = loginPageData{}

// TODO: Переделать на БД
// Закешировать
func getTranslations(t string) (map[string]string, error) {

	var result = make(map[string]string)

	file, err := os.Open(cfg.ROOT_PATH + "/web/ui/translations/ru_RU/loginpage")
	if err != nil {
		return nil, err
	}

	scaner := bufio.NewScanner(file)

	for scaner.Scan() {
		values := strings.Split(scaner.Text(), "=")
		result[values[0]] = values[1]
	}

	return result, nil
}

// Страница логина
func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		var rootPath = cfg.ROOT_PATH + "/web/ui/"

		files := []string{
			rootPath + "login.page.gohtml",
			rootPath + "base.layout.gohtml",
			rootPath + "footers/footer.partial.gohtml",
			rootPath + "headers/headerNoAuth.partial.gohtml",
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

		pageData.Title = "Login--"
		pageData.Translation, err = getTranslations("loginpage")

		if err != nil {
			println(err.Error())
		}

		err = ts.Execute(w, pageData)
		if err != nil {
			log.Print(err)
			http.Error(w, "Internal Server Error", 500)
			return
		}

	} else if r.Method == "POST" {

		var v User
		err := json.NewDecoder(r.Body).Decode(&v)

		w.Header().Set("Content-Type", "application/json")

		if err != nil {
			log.Printf("Invalid data: %s", err)
			json.NewEncoder(w).Encode(response{
				Status: "error",
				Errors: []string{"Invalid data"}})
			return
		}

		if len(v.Login) > 0 && len(v.Password) > 0 {

			row := conn.QueryRow(
				"SELECT password FROM users WHERE user_name=$1",
				v.Login)

			var user database.Users

			err := row.Scan(&user.Password)

			if err != nil {
				if err == sql.ErrNoRows {
					log.Printf("User is not exists with error: %s", err)
					json.NewEncoder(w).Encode(response{
						Status: "error",
						Errors: []string{"User is not exists"}})
					return
				}
				log.Printf("Error scaning user password from DB response with error: %s", err)
				json.NewEncoder(w).Encode(response{
					Status: "error",
					Errors: []string{"Internal server error"}})
				return
			}

			err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(v.Password))

			if err != nil {
				log.Printf("Invalid password: %s", err)
				json.NewEncoder(w).Encode(response{
					Status: "error",
					Errors: []string{"Invalid password"}})
				return
			} else {

				session, _ := sessionStore.Get(r, "session-key")
				session.Values["authenticated"] = true
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
				Status: "error",
				Errors: []string{"Login or password is empty"}})
			return
		}
	}
}
