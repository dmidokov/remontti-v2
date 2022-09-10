package handlers

import (
	"fmt"
	"net/http"

	"github.com/dmidokov/remontti-v2/config"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
)

// Переменная уровня пакета, используется для
// передачи конфигурации в функции обработчики

type HandlersModel struct {
	DB          *pgx.Conn
	Config      *config.Configuration
	CookieStore *sessions.CookieStore
}

type response struct {
	Status  string   `json:"status"`
	Errors  []string `json:"errors"`
	Message string   `json:"message" `
}

func New(db *pgx.Conn, cfg *config.Configuration, cookieStore *sessions.CookieStore) *HandlersModel {
	return &HandlersModel{
		DB:          db,
		Config:      cfg,
		CookieStore: cookieStore,
	}
}

// Возвращает *mux.Router, c handlerFunction
// А также файловые сервер для выдачи статики css/js/jpg/...
func (hm *HandlersModel) Router() (*mux.Router, error) {

	router := mux.NewRouter()

	router.HandleFunc("/", hm.auth(hm.mainPage)).Methods("GET")

	router.HandleFunc("/login", hm.auth(hm.login)).Methods("GET", "POST")
	router.HandleFunc("/logout", hm.auth(hm.logout)).Methods("GET")

	router.HandleFunc("/home", home).Methods("GET")
	router.HandleFunc("/settings", hm.auth(settings)).Methods("Get")

	router.HandleFunc("/companies", hm.auth(hm.companies)).Methods("GET")
	router.HandleFunc("/companies/add", hm.auth(hm.addCompany)).Methods("POST")

	router.HandleFunc("/registration", registration(router)).Methods("Get")

	fileServer := http.FileServer(http.Dir("./web/"))
	router.Handle("/static/{folder}/{file}", http.StripPrefix("/static", fileServer))
	router.Handle("/favicon.ico", fileServer)

	return router, nil
}

func settings(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello! Request was processed settings")
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	target := "https://" + r.Host + r.URL.Path
	http.Redirect(w, r, target, http.StatusMovedPermanently)
}
