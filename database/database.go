package database

import (
	"database/sql"
	"fmt"
)

func ConnectToDB(host, port, user, password, dbname string) (*sql.DB, error) {

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	return conn, nil

}
