package handlers

import (
	"database/sql"
	"net/http"

	"github.com/antonlindstrom/pgstore"
	"github.com/dmidokov/remontti-v2/config"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// Переменная уровня пакета, используется для
// передачи конфигурации в функции обработчики
var cfg *config.Configuration
var conn *sql.DB
var sessionStore *pgstore.PGStore

// Возвращает *mux.Router, c handlerFunction
// А также файловые сервер для выдачи статики css/js/jpg/...
func Router(con *sql.DB, store *pgstore.PGStore, config *config.Configuration) *mux.Router {

	cfg = config
	conn = con
	sessionStore = store

	router := mux.NewRouter()

	router.HandleFunc("/", auth(mainPage)).Methods("GET")

	router.HandleFunc("/login", auth(login)).Methods("GET", "POST")

	router.HandleFunc("/home", home).Methods("GET")

	router.HandleFunc("/registration", registration(router)).Methods("Get")

	router.HandleFunc("/dbcheck", dbCheck(conn))

	fileServer := http.FileServer(http.Dir("./web/"))
	router.Handle("/static/{folder}/{file}", http.StripPrefix("/static", fileServer))
	router.Handle("/favicon.ico", fileServer)

	return router
}
