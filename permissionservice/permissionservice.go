package permissionservice

import (
	"context"
	"errors"
	"github.com/dmidokov/remontti-v2/userservice"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"time"
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
	DB *pgxpool.Pool
}

type Permissions struct {
	PermissionId  int
	ComponentId   int
	ComponentType string
	UserId        int
	Actions       int
	EditTime      int
}

type GroupsPermissions struct {
	PermissionId  int
	ComponentId   int
	ComponentType string
	GroupId       int
	Actions       int
	EditTime      int
}

type Group struct {
	GroupId   int
	GroupName string
}

// ErrPermissionAlreadyExists Ошибка - разрешение уже существует
var ErrPermissionAlreadyExists = errors.New("permissions: Permission already exists")

// Обработка строки ответа от БД
// возвращает структуру Permissions
func permissionsRowProcessing(row pgx.Row) (*Permissions, error) {

	var permission = &Permissions{}

	//TODO:: можно ли принять сразу в запись а не по одному?permission, err := p.GetPermissionsByComponentIdAndUserId(componentId, userId)
	err := row.Scan(&permission.PermissionId, &permission.ComponentId, &permission.ComponentType, &permission.UserId, &permission.Actions, &permission.EditTime)

	if err != nil {
		return nil, err
	}
	return permission, nil

}

func permissionsRowsProcessing(rows pgx.Rows) ([]*Permissions, error) {
	var result []*Permissions

	for rows.Next() {
		var permission = &Permissions{}
		err := rows.Scan(&permission.PermissionId, &permission.ComponentId, &permission.ComponentType, &permission.UserId, &permission.Actions, &permission.EditTime)
		if err != nil {
			log.Print(err)
			continue
		}

		result = append(result, permission)

	}

	return result, nil
}

func groupRowsProcessing(rows pgx.Rows) ([]*Group, error) {
	var result []*Group

	for rows.Next() {
		var group = &Group{}
		err := rows.Scan(&group.GroupId, &group.GroupName)
		if err != nil {
			log.Print(err)
			continue
		}
		result = append(result, group)
	}

	return result, nil
}

func groupRowProcessing(row pgx.Row) (*Group, error) {

	var group = &Group{}

	err := row.Scan(&group.GroupId, &group.GroupName)

	if err != nil {
		println(err.Error())
		return nil, err
	}
	return group, nil

}

// GetPermissionsByComponentIdAndUserId Возвращает Разрешения компонента для пользователя
func (p *PermissionModel) GetPermissionsByComponentIdAndUserId(componentId, userId int, componentType string) (*Permissions, error) {
	sql := "SELECT * FROM remonttiv2.permissions WHERE component_id=$1 AND user_id=$2 AND component_type=$3"
	row := p.DB.QueryRow(context.Background(), sql, componentId, userId, componentType)

	return permissionsRowProcessing(row)
}

// GetPermissionsById Возвращает Разрешения по их идентификатору
func (p *PermissionModel) GetPermissionsById(id int) (*Permissions, error) {
	sql := "SELECT * FROM remonttiv2.permissions WHERE permission_id=$1"
	row := p.DB.QueryRow(context.Background(), sql, id)

	return permissionsRowProcessing(row)
}

func (p *PermissionModel) GetAllPermissions() ([]*Permissions, error) {
	sql := "SELECT * FROM remonttiv2.permissions"
	rows, err := p.DB.Query(context.Background(), sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return permissionsRowsProcessing(rows)
}

func (p *PermissionModel) GetAllGroups() ([]*Group, error) {
	sql := "SELECT * FROM remonttiv2.groups"
	rows, err := p.DB.Query(context.Background(), sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return groupRowsProcessing(rows)
}

// Set Создает новые Разрешения на компонент для пользователя
func (p *PermissionModel) Set(componentId, userId, actions int, componentType string) (*Permissions, error) {
	permission, err := p.GetPermissionsByComponentIdAndUserId(componentId, userId, componentType)

	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}

	if permission == nil {
		sql := "INSERT INTO remonttiv2.permissions (component_id, user_id, actions, edit_time, component_type) VALUES ($1, $2, $3, $4, $5)"
		_, err = p.DB.Exec(context.Background(), sql, componentId, userId, actions, time.Now().Unix(), componentType)

		if err != nil {
			print("ERRORR")
			return nil, err
		}
		print("RETURN OK")
		return p.GetPermissionsByComponentIdAndUserId(componentId, userId, componentType)
	} else {
		print("RETURN ERROR NIL")
		// TODO:: вернуть ошибку о том что запить уже существует
		return nil, err
	}
}

func (p *PermissionModel) Update(componentId, userId, actions int, componentType string) (*Permissions, error) {
	//TODO:: доделать
	permission, err := p.GetPermissionsByComponentIdAndUserId(componentId, userId, componentType)

	print("UPDATE")
	sql := "UPDATE remonttiv2.permissions SET actions=$1, edit_time=$2 WHERE component_id=$3 AND user_id=$4"
	_, err = p.DB.Exec(context.Background(), sql, actions, time.Now().Unix(), componentId, userId)
	if err != nil {
		return nil, err
	}
	return permission, nil
}

func (p *PermissionModel) GetGroupByName(groupName string) (*Group, error) {

	sql := `SELECT * FROM remonttiv2.groups WHERE group_name = $1`

	println(sql)

	row := p.DB.QueryRow(context.Background(), sql, groupName)

	return groupRowProcessing(row)

}

func (p *PermissionModel) AddGroupForUser(userId int, groupName string) error {
	group, err := p.GetGroupByName(groupName)
	if err != nil {
		return err
	}

	sql := `INSERT INTO remonttiv2.users_groups (user_id, group_id) VALUES ($1, $2)`
	_, err = p.DB.Exec(context.Background(), sql, userId, group.GroupId)

	if err != nil {
		return err
	}

	return nil
}

func (p *PermissionModel) GetGroupsByUserId(userId int) ([]int, error) {
	return nil, nil
}

func (p *PermissionModel) AddGroupForUser1(id int, s string) error {
	return nil
}

func (p *PermissionModel) GetPermissionIdByComponentIdAndType(componentId int, componentType string) (int, error) {
	sql := `SELECT * FROM remonttiv2.permissions WHERE component_id = $1 AND component_type=$2`

	row := p.DB.QueryRow(context.Background(), sql, componentId, componentType)

	permissions, err := permissionsRowProcessing(row)
	if err != nil {
		return 0, err
	}

	return permissions.PermissionId, nil
}

func (p *PermissionModel) IsUserHasPermissions(userId int, componentId int, componentType string, actions int) bool {

	var userService = userservice.UserModel{DB: p.DB}
	userGroups, err := userService.GetUserGroups(userId)

	if err != nil {
		return false
	}

	var batch = &pgx.Batch{}

	for _, item := range userGroups {
		sql := `SELECT * FROM remonttiv2.group_permissions WHERE group_id = $1 AND component_id = $2 AND component_type = $3 AND (actions & $4) = $4`

		batch.Queue(sql, item.GroupId, componentId, componentType, actions)
	}

	result := p.DB.SendBatch(context.Background(), batch)
	defer result.Close()

	res, err := result.Query()
	for err == nil {
		if res.Scan() != nil {
			return true
		}
		res, err = result.Query()

	}

	sql := `SELECT * FROM remonttiv2.permissions WHERE user_id = $1 AND component_id=$2 AND component_type = $3 AND  (actions & $4) = $4`

	row := p.DB.QueryRow(context.Background(), sql, userId, componentId, componentType, actions)

	_, err = permissionsRowProcessing(row)
	if err != nil {
		return false
	}

	return true
}

func (p *PermissionModel) IsUserHasPermissions1(userId int, permissionsId int, actions int) bool {

	sql := `SELECT * FROM remonttiv2.permissions WHERE user_id = $1 AND permission_id=$2 AND (actions & $3) = $3`

	row := p.DB.QueryRow(context.Background(), sql, userId, permissionsId, actions)

	_, err := permissionsRowProcessing(row)
	if err != nil {
		return false
	}

	return true
}
