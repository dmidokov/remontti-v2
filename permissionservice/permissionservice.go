package permissionservice

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/jackc/pgx/v4"
)

type ActionsStruct struct {
	VIEW   int
	EDIT   int
	DELETE int
}

var Actions = &ActionsStruct{
	VIEW:   1,
	EDIT:   2,
	DELETE: 4,
}

type PermissionModel struct {
	DB *pgx.Conn
}

type Permissons struct {
	PermissionId  int
	ComponentId   int
	ComponentType string
	UserId        int
	Actions       int
	EditTime      int
}

// ErrPermissionAlreadyExists Ошибка - разрешение уже существует
var ErrPermissionAlreadyExists = errors.New("permissions: Permission already exists")

// Обработка строки ответа от БД
// Возвращает структуру Permissions
func rowProcessing(row pgx.Row) (*Permissons, error) {

	var permission = &Permissons{}

	err := row.Scan(&permission.PermissionId, &permission.ComponentId, &permission.UserId, &permission.Actions, &permission.EditTime)

	if err != nil {
		return nil, err
	}
	return permission, nil

}

func rowsProcessing(rows pgx.Rows) ([]*Permissons, error) {
	var result []*Permissons

	for rows.Next() {
		var permission = &Permissons{}
		err := rows.Scan(&permission.PermissionId, &permission.ComponentId, &permission.ComponentType, &permission.UserId, &permission.Actions, &permission.EditTime)
		if err != nil {
			log.Print(err)
			continue
		}

		result = append(result, permission)

	}

	return result, nil
}

// GetByComponentIdAndUserId Возвращает Разрешения компонента для пользователя
func (p *PermissionModel) GetByComponentIdAndUserId(component_id, user_id int) (*Permissons, error) {
	sql := "SELECT * FROM remonttiv2.permissions WHERE component_id=$1 AND user_id=$2"
	row := p.DB.QueryRow(context.Background(), sql, component_id, user_id)

	return rowProcessing(row)
}

// GetById Возвращает Разрешения по их идентификатору
func (p *PermissionModel) GetById(id int) (*Permissons, error) {
	sql := "SELECT * FROM remonttiv2.permissions WHERE permission_id=$1"
	row := p.DB.QueryRow(context.Background(), sql, id)

	return rowProcessing(row)
}

func (p *PermissionModel) GetAll() ([]*Permissons, error) {
	sql := "SELECT * FROM remonttiv2.permissions WHERE 1=1"
	rows, err := p.DB.Query(context.Background(), sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return rowsProcessing(rows)
}

// Set Создает новые Разрешения на компонент для пользователя
func (p *PermissionModel) Set(componentId, userId, actions int) (*Permissons, error) {

	permission, err := p.GetByComponentIdAndUserId(componentId, userId)

	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}

	if permission != nil {

		sql := "UPDATE remonttiv2.permissions SET actions=$1, edit_time=$2 WHERE component_id=$3 AND user_id=$4"
		_, err = p.DB.Exec(context.Background(), sql, actions, time.Now().Unix(), componentId, userId)
		if err != nil {
			return nil, err
		}
		return permission, nil

	} else {

		sql := "INSERT INTO remonttiv2.permissions (component_id, user_id, actions, edit_time) VALUES ($1, $2, $3, $4)"
		_, err = p.DB.Exec(context.Background(), sql, componentId, userId, actions, time.Now().Unix())

		if err != nil {
			return nil, err
		}
		return p.GetByComponentIdAndUserId(componentId, userId)

	}

}
