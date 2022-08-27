package main

import (
	"log"
	"net/http"

	"github.com/dmidokov/remontti-v2/config"
	"github.com/dmidokov/remontti-v2/database"
	"github.com/dmidokov/remontti-v2/handlers"
	"github.com/dmidokov/remontti-v2/navigationservice"

	"github.com/gorilla/sessions"
)

type application struct {
	config     *config.Configuration
	navigation *navigationservice.NavigationModel
	database   *database.DatabaseModel
	handlers   *handlers.HandlersModel
}

func main() {

	var app application

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

	app = application{
		config:     config,
		navigation: &navigationservice.NavigationModel{DB: conn},
		database:   &database.DatabaseModel{DB: conn},
		handlers:   &handlers.HandlersModel{DB: conn},
	}

	log.Print("Подкотовка БД")
	err = app.database.Prepare(config)
	if err != nil {
		log.Fatalf("Подготовка завершилась с ошибкой: %s", err)
	}

	log.Print("Подготовка хранилища сесссий")
	store := sessions.NewCookieStore([]byte(config.SESSIONS_SECRET))

	// Регистрируем обработчики получаем роутер
	log.Print("Регистрация обработчиком запросов")
	router, err := app.handlers.Router(conn, store, config)
	if err != nil {
		log.Fatalf("Регистрация завершилась с ошибкой: %s", err)
	}

	log.Print("Сервис запущен и готов к приему запросов")

	// Запускаем лиснер
	log.Fatal(http.ListenAndServe(":8000", router))
}
