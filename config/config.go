// Пакет config осуществляет загрузку конфигурации приложения
// из переменных окружения
package config

import (
	"errors"
	"os"
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

// Имена переменных окружения
const (
	DB_USER     = "DB_USER_NAME"
	DB_PASSWORD = "DB_USER_PASSWORD"
	DB_PORT     = "DB_PORT"
	DB_HOST     = "DB_HOST"
	DB_NAME     = "DB_NAME"
	ROOT_PATH   = "ROOT_PATH"
)

// Функция заполняет структуру config из переменных окружения
// Возвращает заполненую структу и ошибку
//
// Примечание: возможно имеет смысл создать слайс
// с именами переменных которые требуются и далее загрузить
// их в цикле без дублирования кода
func LoadConfig() (*Configuration, error) {

	var config = Configuration{}
	var exist bool

	if config.DB_USER, exist = os.LookupEnv(DB_USER); !exist {
		return &Configuration{}, errors.New("database  username is missing")
	}

	if config.DB_PASSWORD, exist = os.LookupEnv(DB_PASSWORD); !exist {
		return &Configuration{}, errors.New("database password is missing")
	}

	if config.DB_HOST, exist = os.LookupEnv(DB_HOST); !exist {
		return &Configuration{}, errors.New("database host is missing")
	}

	if config.DB_PORT, exist = os.LookupEnv(DB_PORT); !exist {
		return &Configuration{}, errors.New("database port is missing")
	}

	if config.DB_NAME, exist = os.LookupEnv(DB_NAME); !exist {
		return &Configuration{}, errors.New("database name is missing")
	}

	if config.ROOT_PATH, exist = os.LookupEnv(ROOT_PATH); !exist {
		return &Configuration{}, errors.New("application root path is empty")
	}

	// Возвращаем конфиг если не вышли ранее с ошибкой
	return &config, nil
}
