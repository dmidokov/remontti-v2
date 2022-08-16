package database

import (
	"database/sql"
	"fmt"

	"github.com/dmidokov/remontti-v2/config"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

func ConnectToDB(host, port, user, password, dbname string) (*sql.DB, error) {

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	db = conn

	return conn, nil

}

func PrepareDB(config *config.Configuration) error {

	sql := fmt.Sprintf(CreateUsersTable, config.DB_USER)
	_, err := db.Exec(sql)

	if err != nil {
		return err
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(config.ADMIN_PASSWORD), 14)

	if err != nil {
		return err
	}

	var count int
	db.QueryRow(SelectAdminCount).Scan(&count)

	if count == 0 {
		sql = fmt.Sprintf(InsertAdminUser, string(bytes))
		_, err = db.Exec(sql)

		if err != nil {
			return err
		}
	}

	return nil
}
