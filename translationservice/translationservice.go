package translationservice

import (
	"context"
	"fmt"
	"github.com/dmidokov/remontti-v2/config"
	"github.com/jackc/pgx/v4"
	"log"
	"time"
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

type TranslationsMemCache struct {
	Time   int64
	Result []*Translation
}

var MemCache = map[string]*TranslationsMemCache{}

func (t *TranslationsModel) Push(pagenames string, sliceOfTranslations []*Translation) {
	MemCache[pagenames] = &TranslationsMemCache{
		Time:   time.Now().Unix(),
		Result: sliceOfTranslations,
	}
}
func (t *TranslationsModel) Pop(pagenames string) *TranslationsMemCache {
	if v, b := MemCache[pagenames]; b == true && time.Now().Unix()-v.Time < 10 {
		return MemCache[pagenames]
	} else {
		return nil
	}
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
