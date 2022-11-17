package translationservice

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dmidokov/remontti-v2/config"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"time"
)

type TranslationsModel struct {
	DB     *pgxpool.Pool
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

const CACHE_TTL = 5

type TranslationsMemCache struct {
	EditTime int64
	Result   []*Translation
}
type MemCacheKey uint

var memCache = map[MemCacheKey]*TranslationsMemCache{}

func (t *TranslationsModel) Push(pageNames string, sliceOfTranslations []*Translation) {
	hash := hash(pageNames)
	memCache[hash] = &TranslationsMemCache{
		EditTime: time.Now().Unix() + CACHE_TTL,
		Result:   sliceOfTranslations,
	}
}
func (t *TranslationsModel) Pop(pageNames string) *TranslationsMemCache {
	hash := hash(pageNames)
	if v, b := memCache[hash]; b == true && time.Now().Unix() < v.EditTime {
		return memCache[hash]
	} else {
		return nil
	}
}

func cachePrint(pushOrPop string) {
	println("======== Cache ", pushOrPop, time.Now().Unix(), "=========")
	for key, value := range memCache {
		v, _ := json.Marshal(value)
		println(key, "=>", string(v))
	}
}

// Get Закешировать
func (t *TranslationsModel) Get(pageNames ...string) ([]*Translation, error) {

	sql := `SELECT * FROM remonttiv2.translations WHERE `
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
