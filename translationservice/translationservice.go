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

const CACHE_TTL = 60

type TranslationsMemCache struct {
	Time   int64
	Result []*Translation
}
type MemCacheKey uint

var MemCache = map[MemCacheKey]*TranslationsMemCache{}

func (t *TranslationsModel) Push(pageNames string, sliceOfTranslations []*Translation) {
	hash := hash(pageNames)
	MemCache[hash] = &TranslationsMemCache{
		Time:   time.Now().Unix(),
		Result: sliceOfTranslations,
	}
}
func (t *TranslationsModel) Pop(pageNames string) *TranslationsMemCache {
	hash := hash(pageNames)
	if v, b := MemCache[hash]; b == true && time.Now().Unix()-v.Time < CACHE_TTL {
		return MemCache[hash]
	} else {
		return nil
	}
}

// Get Закешировать
func (t *TranslationsModel) Get(pageNames ...string) ([]*Translation, error) {

	sql := `SELECT * FROM remonttiv2.translations WHERE true`
	for i, name := range pageNames {
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

	sql := `SELECT * FROM remonttiv2.translations WHERE true`

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

func hash(str string) MemCacheKey {
	var k MemCacheKey = 67
	var mod MemCacheKey = 1e9 + 7
	var h MemCacheKey = 0
	var m MemCacheKey = 1

	for _, v := range str {
		var x = MemCacheKey(v - 96)
		h = (h + m*x) % mod
		m = (h * k) % mod
	}
	return h
}
