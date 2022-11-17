package companyservice

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"time"

	"github.com/jackc/pgx/v4"
)

type CompanyModel struct {
	DB *pgxpool.Pool
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

func (c *CompanyModel) GetCompanyByName(companyName string) (*Company, error) {

	sql := `SELECT * FROM remonttiv2.companies WHERE company_name=$1;`

	row := c.DB.QueryRow(context.Background(), sql, companyName)

	return rowProcessing(row)

}

func (c *CompanyModel) GetCompanyByHostName(hostName string) (*Company, error) {

	sql := `SELECT * FROM remonttiv2.companies WHERE host_name=$1;`

	row := c.DB.QueryRow(context.Background(), sql, hostName)

	return rowProcessing(row)

}

func rowProcessing(row pgx.Row) (*Company, error) {

	var company = &Company{}

	err := row.Scan(&company.CompanyId, &company.CompanyName, &company.HostName, &company.EditTime)

	if err != nil {
		return nil, err
	}
	return company, nil

}

func (c *CompanyModel) Add(name, host string) (*Company, error) {

	time := time.Now().Unix()

	sql := "INSERT INTO remonttiv2.companies (company_name, host_name, edit_time) VALUES($1, $2, $3)"

	_, err := c.DB.Exec(context.Background(), sql, name, host, time)
	if err != nil {
		return nil, err
	}

	sql = "SELECT * FROM remonttiv2.companies WHERE company_name=$1 AND host_name=$2 AND edit_time=$3"

	row := c.DB.QueryRow(context.Background(), sql, name, host, time)

	return rowProcessing(row)
}
