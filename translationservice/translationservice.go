package translationservice

import (
	"context"
	"fmt"
	"log"

	"github.com/dmidokov/remontti-v2/config"
	"github.com/jackc/pgx/v4"
)

type TranslationsModel struct {
	DB     *pgx.Conn
	Config *config.Configuration // Возможно лишнее!!!
}

type Translation struct {
	Id       int
	Name     string
	Label    string
	Ru       string
	En       string
	EditTime int
}

// Get Закешировать
func (t *TranslationsModel) Get(pagenames ...string) ([]*Translation, error) {

	sql := `SELECT * FROM remonttiv2.translations WHERE `
	for i, name := range pagenames {
		if i > 0 {
			sql += " OR "
		}
		sql += fmt.Sprintf("name='%s'", name)
	}

	rows, err := t.DB.Query(context.Background(), sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return rowsProcessing(rows)
}

func (t *TranslationsModel) GetAll() ([]*Translation, error) {

	sql := `SELECT * FROM remonttiv2.translations WHERE 1=1`

	rows, err := t.DB.Query(context.Background(), sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return rowsProcessing(rows)
}

func rowsProcessing(rows pgx.Rows) ([]*Translation, error) {
	var result []*Translation

	for rows.Next() {
		var translation = &Translation{}
		err := rows.Scan(&translation.Id, &translation.Name, &translation.Label, &translation.Ru, &translation.En, &translation.EditTime)
		if err != nil {
			log.Print(err)
			continue
		}

		result = append(result, translation)

	}

	return result, nil
}
