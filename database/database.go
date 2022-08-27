package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dmidokov/remontti-v2/config"
	"github.com/dmidokov/remontti-v2/navigationservice"
	"github.com/dmidokov/remontti-v2/permissionservice"
	"github.com/dmidokov/remontti-v2/translationservice"
	"github.com/dmidokov/remontti-v2/userservice"
	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
)

var db *pgx.Conn

type DatabaseModel struct {
	DB *pgx.Conn
}

func ConnectToDB(dbHost, dbPort, dbUser, dbPassword, dbName string) (*pgx.Conn, error) {

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	conn, err := pgx.Connect(context.Background(), psqlInfo)
	if err != nil {
		return nil, err
	}

	err = conn.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	db = conn

	return conn, nil

}

// Подготавка БД к работе
func (pg *DatabaseModel) Prepare(cfg *config.Configuration) error {

	log.Print("Создание таблиц")
	err := pg.createTables(cfg)
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

	sqlInsert := `INSERT INTO public.users 
				(company_id, user_name, password, last_login_date, last_login_error_date) 
			VALUES 
				($1, $2, $3, $4, $5);`
	sqlUpdate := `UPDATE public.users SET password=$1 WHERE user_name=$2 AND company_id=$3`

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

	sqlInsert := `INSERT INTO public.navigation (item_type, link, label, edit_time) VALUES ($1, $2, $3, $4)`
	sqlUpdate := `UPDATE public.navigation SET item_type=$1, label=$2, edit_time=$3 WHERE link=$4`

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

	sqlInsert := `INSERT INTO public.permissions (component_id, user_id, actions, edit_time) VALUES ($1, $2, $3, $4)`
	sqlUpdate := `UPDATE public.permissions SET actions=$1, edit_time=$2 WHERE component_id=$3 AND user_id=$4`

	for _, itemData := range permissionsData {

		var itemExist bool = false
		for _, permission := range permissions {
			if permission.UserIid == itemData.UserIid && permission.ComponentId == itemData.ComponentId {
				itemExist = true
				break
			}
		}

		if itemExist {
			batch.Queue(sqlUpdate, itemData.Actions, time.Now().Unix(), itemData.ComponentId, itemData.UserIid)
		} else {
			batch.Queue(sqlInsert, itemData.ComponentId, itemData.UserIid, itemData.Actions, itemData.EditTime)
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
	sqlInsert := `INSERT INTO public.translations (name, label, ru, en, edit_time) VALUES ($1, $2, $3, $4, $5)`
	sqlUpdate := `UPDATE public.translations SET ru=$1, en=$2, edit_time=$3 WHERE name=$4 AND label=$5`

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
