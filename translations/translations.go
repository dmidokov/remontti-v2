package translations

import (
	"bufio"
	"os"
	"strings"

	"github.com/dmidokov/remontti-v2/config"
)

// TODO: Переделать на БД
// Закешировать
func GetTranslations(fileName string, cfg *config.Configuration) (map[string]string, error) {

	var result = make(map[string]string)

	file, err := os.Open(cfg.ROOT_PATH + "/web/ui/translations/ru_RU/" + fileName)
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
