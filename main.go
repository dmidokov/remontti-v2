package main

import (
	"log"
	"net/http"

	"github.com/dmidokov/remontti-v2/config"
	"github.com/dmidokov/remontti-v2/database"
	"github.com/dmidokov/remontti-v2/handlers"

	"github.com/gorilla/sessions"
)

func main() {

	log.Print("Starting the service...")

	// Пытаемся загрузить конфигурацию
	// Если нет выходим с ошибкой
	log.Print("Trying to load configuration")
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Подключаемся в БД, параметры подключения берем из конфигурации
	// Если нет выходим с ошибкой
	log.Print("Trying to connect to database")
	conn, err := database.ConnectToDB(
		config.DB_HOST,
		config.DB_PORT,
		config.DB_USER,
		config.DB_PASSWORD,
		config.DB_NAME)
	if err != nil {
		log.Fatalf("Connecting to DB finish with error : %s", err)
	}

	err = database.PrepareDB(config)
	if err != nil {
		log.Fatalf("Databas preparing ending with error: %s", err)
	}

	log.Print("Prepare sessions storage")
	store := sessions.NewCookieStore([]byte(config.SESSIONS_SECRET))

	// Регистрируем обработчики получаем роутер
	log.Print("Registrate handlers")
	router := handlers.Router(conn, store, config)

	log.Print("The service is ready to listen and serve")

	// Запускаем лиснер
	log.Fatal(http.ListenAndServe(":8000", router))
}
