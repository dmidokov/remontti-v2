package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"time"

	"github.com/dmidokov/remontti-v2/companyservice"
	"github.com/dmidokov/remontti-v2/config"
	"github.com/dmidokov/remontti-v2/navigationservice"
	"github.com/dmidokov/remontti-v2/permissionservice"
	"github.com/dmidokov/remontti-v2/translationservice"
	"github.com/dmidokov/remontti-v2/userservice"
	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
)

type DatabaseModel struct {
	DB     *pgxpool.Pool
	Logger *logrus.Logger
}

func ConnectToDB(dbHost, dbPort, dbUser, dbPassword, dbName string) (*pgxpool.Pool, error) {

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
	conn, err := pgxpool.Connect(context.Background(), psqlInfo)
	//conn, err := pgx.Connect(context.Background(), psqlInfo)
	if err != nil {
		return nil, err
	}

	err = conn.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return conn, nil

}

// Подготавка БД к работе
func (pg *DatabaseModel) Prepare(cfg *config.Configuration) error {

	var log = pg.Logger

	if cfg.DELETE_TABLES_BEFORE_START == 1 {
		log.Print("Пересоздание схемы")
		sql := `
			DROP SCHEMA remonttiv2 CASCADE;
			CREATE SCHEMA remonttiv2;`
		_, err := pg.DB.Exec(context.Background(), sql)
		if err != nil {
			return err
		}
	}

	log.Print("Создание таблиц")
	err := pg.createTables(cfg)
	if err != nil {
		return err
	}

	log.Print("Создание компаний")
	err = pg.insertCompaniesData(cfg)
	if err != nil {
		return err
	}

	log.Print("Создание пользователей")
	err = pg.insertUsersData(cfg)
	if err != nil {
		return err
	}

	log.Print("Создание пунктов меню")
	err = pg.insertNavigationData(cfg)
	if err != nil {
		return err
	}

	log.Print("Создание пермишенов")
	err = pg.insertPermissionsData(cfg)
	if err != nil {
		return err
	}

	log.Print("Создание переводов")
	err = pg.insertTranslationsData(cfg)
	if err != nil {
		return err
	}

	return nil
}

func (pg *DatabaseModel) createTables(cfg *config.Configuration) error {

	var batch *pgx.Batch = &pgx.Batch{}

	batch.Queue(CreateUsersTableSQL)
	batch.Queue(CreateNavigationTableSQL)
	batch.Queue(CreatePermissionsTableSQL)
	batch.Queue(CreateTranslationsTableSQL)
	batch.Queue(CreateCompaniesTableSQL)

	result := pg.DB.SendBatch(context.Background(), batch)
	defer result.Close()

	var err error

	_, err = result.Exec()
	for err == nil {
		_, err = result.Exec()
	}

	if err.Error() == "no result" {
		return nil
	} else {
		return err
	}
}

// Создает таблицу пользователей если нет
func (pg *DatabaseModel) insertUsersData(cfg *config.Configuration) error {

	var userservice userservice.UserModel = userservice.UserModel{DB: pg.DB}

	var batch *pgx.Batch = &pgx.Batch{}

	usersData := GetUserDataToInsert(cfg)

	allUsers, err := userservice.GetAll()
	if err != nil {
		return err
	}

	sqlInsert := `INSERT INTO remonttiv2.users 
				(company_id, user_name, password, last_login_date, last_login_error_date) 
			VALUES 
				($1, $2, $3, $4, $5);`
	sqlUpdate := `UPDATE remonttiv2.users SET password=$1 WHERE user_name=$2 AND company_id=$3`

	for _, userData := range usersData {

		bytes, err := bcrypt.GenerateFromPassword([]byte(userData.Password), 14)
		if err != nil {
			return err
		}

		userExist := false
		for _, userInTable := range allUsers {
			if userInTable.CompanyId == userData.CompanyId && userInTable.UserName == userData.UserName {
				userExist = true
				break
			}
		}

		if userExist {
			batch.Queue(sqlUpdate, string(bytes), userData.UserName, userData.CompanyId)
		} else {
			batch.Queue(sqlInsert, userData.CompanyId, userData.UserName, string(bytes), 0, 0)
		}

	}

	result := pg.DB.SendBatch(context.Background(), batch)
	defer result.Close()

	for err == nil {
		_, err = result.Exec()
	}

	if err.Error() == "no result" {
		return nil
	} else {
		return err
	}
}

func (pg *DatabaseModel) insertNavigationData(cfg *config.Configuration) error {
	var navigationservice navigationservice.NavigationModel = navigationservice.NavigationModel{DB: pg.DB}

	var batch *pgx.Batch = &pgx.Batch{}

	navigationData := GetNavigationDataToInsert(cfg)

	allItems, err := navigationservice.GetAll()
	if err != nil {
		return err
	}

	sqlInsert := `INSERT INTO remonttiv2.navigation (item_type, link, label, edit_time) VALUES ($1, $2, $3, $4)`
	sqlUpdate := `UPDATE remonttiv2.navigation SET item_type=$1, label=$2, edit_time=$3 WHERE link=$4`

	for _, itemData := range navigationData {

		var itemExist bool = false
		for _, itemInTable := range allItems {
			if itemInTable.Item_type == itemData.Item_type && itemInTable.Label == itemData.Label {
				itemExist = true
				break
			}
		}

		if itemExist {
			batch.Queue(sqlUpdate, itemData.Item_type, itemData.Label, itemData.EditTime, itemData.Link)
		} else {
			batch.Queue(sqlInsert, itemData.Item_type, itemData.Link, itemData.Label, itemData.EditTime)
		}

	}

	result := pg.DB.SendBatch(context.Background(), batch)
	defer result.Close()

	for err == nil {
		_, err = result.Exec()
	}

	if err.Error() == "no result" {
		return nil
	} else {
		return err
	}

}

func (pg *DatabaseModel) insertPermissionsData(cfg *config.Configuration) error {

	var permissionservice permissionservice.PermissionModel = permissionservice.PermissionModel{DB: pg.DB}

	var batch *pgx.Batch = &pgx.Batch{}

	permissionsData, err := pg.GetPermissionsDataToInsert(cfg)
	if err != nil {
		return err
	}

	permissions, err := permissionservice.GetAll()
	if err != nil {
		return err
	}

	sqlInsert := `INSERT INTO remonttiv2.permissions (component_id, user_id, actions, edit_time, component_type) VALUES ($1, $2, $3, $4, $5)`
	sqlUpdate := `UPDATE remonttiv2.permissions SET actions=$1, edit_time=$2 WHERE component_id=$3 AND user_id=$4 AND component_type=$5`

	for _, itemData := range permissionsData {

		var itemExist bool = false
		for _, permission := range permissions {
			if permission.UserId == itemData.UserId && permission.ComponentId == itemData.ComponentId {
				itemExist = true
				break
			}
		}

		if itemExist {
			batch.Queue(sqlUpdate, itemData.Actions, time.Now().Unix(), itemData.ComponentId, itemData.UserId, itemData.ComponentType)
		} else {
			batch.Queue(sqlInsert, itemData.ComponentId, itemData.UserId, itemData.Actions, itemData.EditTime, itemData.ComponentType)
		}

	}

	result := pg.DB.SendBatch(context.Background(), batch)
	defer result.Close()

	for err == nil {
		_, err = result.Exec()
	}

	if err.Error() == "no result" {
		return nil
	} else {
		return err
	}

}

func (pg *DatabaseModel) insertTranslationsData(cfg *config.Configuration) error {

	translationservice := translationservice.TranslationsModel{DB: pg.DB}

	var batch *pgx.Batch = &pgx.Batch{}

	translationsData := GetTranslationsDataToInsert(cfg)

	translations, err := translationservice.GetAll()
	if err != nil {
		return err
	}

	// TODO: make the correct SQL requests
	sqlInsert := `INSERT INTO remonttiv2.translations (name, label, ru, en, edit_time) VALUES ($1, $2, $3, $4, $5)`
	sqlUpdate := `UPDATE remonttiv2.translations SET ru=$1, en=$2, edit_time=$3 WHERE name=$4 AND label=$5`

	for _, itemData := range translationsData {

		var itemExist bool = false
		for _, translation := range translations {
			if translation.Name == itemData.Name && translation.Label == itemData.Label {
				itemExist = true
				break
			}
		}

		if itemExist {
			batch.Queue(sqlUpdate, itemData.Ru, itemData.En, itemData.EditTime, itemData.Name, itemData.Label)

		} else {
			batch.Queue(sqlInsert, itemData.Name, itemData.Label, itemData.Ru, itemData.En, itemData.EditTime)
		}

	}

	result := pg.DB.SendBatch(context.Background(), batch)
	defer result.Close()

	for err == nil {
		_, err = result.Exec()
	}

	if err.Error() == "no result" {
		return nil
	} else {
		println(err.Error())
		return err
	}

}

func (pg *DatabaseModel) insertCompaniesData(cfg *config.Configuration) error {

	var companyservice companyservice.CompanyModel = companyservice.CompanyModel{DB: pg.DB}

	var batch *pgx.Batch = &pgx.Batch{}

	companiesData := GetCompaniesDataToInsert()

	allCompanies, err := companyservice.GetAll()
	if err != nil {
		return err
	}

	sqlInsert := `INSERT INTO remonttiv2.companies 
					(company_id, company_name, host_name, edit_time) 
				VALUES 
				($1, $2, $3, $4);`
	sqlUpdate := `UPDATE remonttiv2.companies SET host_name=$1, edit_time=$2 WHERE company_id=$3 AND company_name=$4`

	for _, companyData := range companiesData {

		companyExist := false
		for _, companyInTable := range allCompanies {
			if companyInTable.CompanyId == companyData.CompanyId && companyInTable.CompanyName == companyData.CompanyName {
				companyExist = true
				break
			}
		}

		if companyExist {
			batch.Queue(sqlUpdate, companyData.HostName, companyData.EditTime, companyData.CompanyId, companyData.CompanyName)
		} else {
			batch.Queue(sqlInsert, companyData.CompanyId, companyData.CompanyName, companyData.HostName, companyData.EditTime)
		}

	}

	result := pg.DB.SendBatch(context.Background(), batch)
	defer result.Close()

	for err == nil {
		_, err = result.Exec()
	}

	if err.Error() == "no result" {
		return nil
	} else {
		return err
	}
}
