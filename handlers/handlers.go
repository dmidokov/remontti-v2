package handlers

import (
	"database/sql"
	"net/http"

	"github.com/RemonttiCRM/remontti-v2/config"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// Переменная уровня пакета, используется для 
// передачи конфигурации в функции обработчики
// 
// Возмодно стоит отпраделить структуру под конфиг в текущем пакете
// и обновить обработчики сдав их методами этой структуры  
var cfg *config.Configuration

// Возвращает *mux.Router, c handlerFunction
// А также файловые сервер для выдачи статики css/js/jpg/...
func Router(conn *sql.DB, config *config.Configuration) *mux.Router {

	cfg = config

	router := mux.NewRouter()

	router.HandleFunc("/", mainPage).Methods("GET")

	router.HandleFunc("/home", home).Methods("GET")

	router.HandleFunc("/registration", registration(router)).Methods("Get")

	router.HandleFunc("/dbcheck", dbCheck(conn))

	fileServer := http.FileServer(http.Dir("./static/"))
	router.Handle("/static/{folder}/{file}", http.StripPrefix("/static", fileServer))

	return router
}