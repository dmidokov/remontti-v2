package config

// Имена переменных окружения
const (
	DB_USER     = "DB_USER_NAME"
	DB_PASSWORD = "DB_USER_PASSWORD"
	DB_PORT     = "DB_PORT"
	DB_HOST     = "DB_HOST"
	DB_NAME     = "DB_NAME"
	ROOT_PATH   = "ROOT_PATH"
)

// Структура для хранения конфигруации
type Configuration struct {
	DB_USER     string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
	ROOT_PATH   string
}
