package companyservice

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

type CompanyModel struct {
	DB *pgx.Conn
}

type Company struct {
	CompanyId   int
	CompanyName string
	HostName    string
	EditTime    int
}

func rowsProcessing(rows pgx.Rows) ([]*Company, error) {
	var result []*Company

	for rows.Next() {
		var item = &Company{}
		err := rows.Scan(&item.CompanyId, &item.CompanyName, &item.HostName, &item.EditTime)
		if err != nil {
			log.Print(err)
			continue
		}

		result = append(result, item)

	}

	return result, nil
}

func (c *CompanyModel) GetAll() ([]*Company, error) {
	sql := `SELECT * FROM remonttiv2.companies WHERE 1=1;`

	rows, err := c.DB.Query(context.Background(), sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return rowsProcessing(rows)
}
