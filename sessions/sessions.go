package sessions

import (
	"fmt"

	"github.com/antonlindstrom/pgstore"
)

func GetStore(host, port, user, password, dbname string) (*pgstore.PGStore, error) {

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	store, err := pgstore.NewPGStore(connectionString, []byte("secret-key"))

	if err != nil {
		return nil, err
	}

	return store, nil

}
