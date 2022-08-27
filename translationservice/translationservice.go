package translationservice

import (
	"context"
	"log"

	"github.com/dmidokov/remontti-v2/config"
	"github.com/jackc/pgx/v4"
)

type TranslationsModel struct {
	DB *pgx.Conn
}

type Translation struct {
	Id       int
	Name     string
	Label    string
	Ru       string
	En       string
	EditTime int
}

// TODO: Переделать на БД
// Закешировать
func (t *TranslationsModel) Get(pagename string, cfg *config.Configuration) ([]*Translation, error) {

	sql := `SELECT * FROM public.translations WHERE name=$1`

	rows, err := t.DB.Query(context.Background(), sql, pagename)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return rowsProcessing(rows)
}

func (t *TranslationsModel) GetAll() ([]*Translation, error) {

	sql := `SELECT * FROM public.translations WHERE 1=1`

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
