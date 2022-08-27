package handlers

import (
	"fmt"
	"net/http"

	"github.com/dmidokov/remontti-v2/config"
	"github.com/dmidokov/remontti-v2/navigationservice"
	"github.com/dmidokov/remontti-v2/translationservice"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
)

// Переменная уровня пакета, используется для
// передачи конфигурации в функции обработчики
var cfg *config.Configuration
var conn *pgx.Conn
var sessionStore *sessions.CookieStore

var navigation navigationservice.NavigationModel = navigationservice.NavigationModel{}
var translation translationservice.TranslationsModel = translationservice.TranslationsModel{}

type HandlersModel struct {
	DB *pgx.Conn
}

// Возвращает *mux.Router, c handlerFunction
// А также файловые сервер для выдачи статики css/js/jpg/...
func (hm *HandlersModel) Router(con *pgx.Conn, store *sessions.CookieStore, config *config.Configuration) (*mux.Router, error) {

	cfg = config
	conn = con
	sessionStore = store

	navigation.DB = con
	translation.DB = con

	router := mux.NewRouter()

	router.HandleFunc("/", auth(mainPage)).Methods("GET")

	router.HandleFunc("/login", auth(login)).Methods("GET", "POST")
	router.HandleFunc("/logout", auth(logout)).Methods("GET")

	router.HandleFunc("/home", home).Methods("GET")

	router.HandleFunc("/registration", registration(router)).Methods("Get")

	fileServer := http.FileServer(http.Dir("./web/"))
	router.Handle("/static/{folder}/{file}", http.StripPrefix("/static", fileServer))
	router.Handle("/favicon.ico", fileServer)

	items, err := navigation.GetAll()
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		switch item.Link {
		case "/settings":
			router.HandleFunc(item.Link, auth(settings))
		case "/logout":
			router.HandleFunc(item.Link, auth(logout1))
		}
	}

	return router, nil
}

func settings(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello! Request was processed settings")
}
func logout1(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello! Request was processed on logout1")
}
