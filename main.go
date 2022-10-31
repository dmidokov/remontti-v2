package main

import (
	"errors"
	"github.com/dmidokov/remontti-v2/config"
	"github.com/dmidokov/remontti-v2/database"
	"github.com/dmidokov/remontti-v2/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

func main() {

	finish := make(chan bool)

	log.Print("Запуск сервиса...")

	// Пытаемся загрузить конфигурацию
	// Если нет выходим с ошибкой
	log.Print("Загрузка конфигурации")
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Подключаемся в БД, параметры подключения берем из конфигурации
	// Если нет выходим с ошибкой
	log.Print("Подключение к БД")
	conn, err := database.ConnectToDB(
		config.DB_HOST,
		config.DB_PORT,
		config.DB_USER,
		config.DB_PASSWORD,
		config.DB_NAME)
	if err != nil {
		log.Fatalf("Подключение завершилось с ошибкой : %s", err)
	}

	log.Print("Подготовка хранилища сесссий")
	store := sessions.NewCookieStore([]byte(config.SESSIONS_SECRET))

	handler := handlers.New(conn, config, store)
	database := &database.DatabaseModel{DB: conn}

	log.Print("Подкотовка БД")
	err = database.Prepare(config)
	if err != nil {
		log.Fatalf("Подготовка завершилась с ошибкой: %s", err)
	}

	// Регистрируем обработчики получаем роутер
	log.Print("Регистрация обработчиком запросов")
	var router *mux.Router

	if config.MODE == "production" {
		router, err = handler.Router(false)
	} else if config.MODE == "dev" {
		router, err = handler.Router(true)
	} else {
		router, err = nil, errors.New("Неизвестный режим запуска")
	}

	if err != nil {
		log.Fatalf("Регистрация завершилась с ошибкой: %s", err)
	}

	log.Print("Сервис запущен и готов к приему запросов")

	if err != nil {
		log.Fatalf("Регистрация завершилась с ошибкой: %s", err)
	}

	// Запускаем лиснеры
	go func() {
		log.Fatal(http.ListenAndServeTLS(":443", "secrets/localhost2.crt", "secrets/localhost2.key", router))
	}()

	go func() {
		log.Fatal(http.ListenAndServe(":80", http.HandlerFunc(handlers.Redirect)))
	}()

	<-finish
}
