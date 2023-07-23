package companyservice

import (
	"context"
	"github.com/dmidokov/remontti-v2/permissionservice"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"log"
	"time"

	"github.com/jackc/pgx/v4"
)

type CompanyModel struct {
	DB     *pgxpool.Pool
	Logger *logrus.Logger
}

type Company struct {
	CompanyId   int
	CompanyName string
	HostName    string
	EditTime    int
}

type CompaniesResult struct {
	CompanyName string `json:"company"`
	HostName    string `json:"host"`
	CompanyId   int    `json:"id"`
}

type CompanyNameResponse struct {
	CompanyName string `json:"company_name"`
}

const ComponentType = "company"

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
	rows, err := c.DB.Query(context.Background(), getAll)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return rowsProcessing(rows)
}

func (c *CompanyModel) GetCompanyByName(companyName string) (*Company, error) {
	row := c.DB.QueryRow(context.Background(), getCompanyByName, companyName)
	return rowProcessing(row)
}

func (c *CompanyModel) GetCompanyById(companyId int) (*Company, error) {
	row := c.DB.QueryRow(context.Background(), getCompanyById, companyId)
	return rowProcessing(row)
}

func (c *CompanyModel) GetCompanyByHostName(hostName string) (*Company, error) {
	row := c.DB.QueryRow(context.Background(), getCompanyByHostName, hostName)
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

	_, err := c.DB.Exec(context.Background(), insertCompany, name, host, time)
	if err != nil {
		return nil, err
	}

	row := c.DB.QueryRow(context.Background(), getCompanyByNameHostTime, name, host, time)

	return rowProcessing(row)
}

// Delete удаляет компанию, а также запись в таблице пермишенов, с появлением таблицы групповых разрешений необходимо дополнить функцию
func (c *CompanyModel) Delete(companyId int) (*Company, error) {

	company, err := c.GetCompanyById(companyId)

	if err != nil {
		return nil, err
	}

	_, err = c.DB.Exec(context.Background(), deleteByCompanyId, companyId)
	if err != nil {
		return nil, err
	}

	_, err = c.DB.Exec(context.Background(), deletePermissionByComponentIdAndType, companyId, ComponentType)
	if err != nil {
		return nil, err
	}

	return company, nil
}

func (c *CompanyModel) GetAllForUser(userId int) ([]*CompaniesResult, error) {

	rows, err := c.DB.Query(context.Background(), getCompaniesForUser, permissionservice.Actions.VIEW, userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	processedRows, _ := rowsProcessing(rows)

	result := make([]*CompaniesResult, 0)
	for _, item := range processedRows {
		result = append(
			result,
			&CompaniesResult{CompanyName: item.CompanyName, HostName: item.HostName, CompanyId: item.CompanyId},
		)
	}

	return result, nil
}

func (c *CompanyModel) GetUserCompanyName(userId int) string {

	row := c.DB.QueryRow(context.Background(), getUserCompanyName, userId)

	var result = &CompanyNameResponse{}
	err := row.Scan(&result.CompanyName)
	if err != nil {
		println(err.Error())
		return "-"
	}

	return result.CompanyName
}

func (c *CompanyModel) Update(id int, name, host string) (*Company, error) {

	_, err := c.DB.Exec(context.Background(), updateCompany, name, host, id)
	if err != nil {
		return nil, err
	}

	row := c.DB.QueryRow(context.Background(), getCompanyById, id)

	return rowProcessing(row)
}
