package navigationservice

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"time"

	"github.com/dmidokov/remontti-v2/permissionservice"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type NavigationItem struct {
	Id        int
	Item_type int
	Link      string
	Label     string
	EditTime  int64
}

type NavigationItemInsert struct {
	ItemType int
	Link     string
	Label    string
}

type NavigationModel struct {
	DB *pgxpool.Pool
}

var ErrItemAlreadyExists = errors.New("navigation: Item already exists")

func New(db *pgxpool.Pool) *NavigationModel {
	return &NavigationModel{
		DB: db,
	}
}

func rowProcessing(row pgx.Row) (*NavigationItem, error) {

	var navigation = &NavigationItem{}

	err := row.Scan(&navigation.Id, &navigation.Item_type, &navigation.Link, &navigation.Label, &navigation.EditTime)

	if err != nil {
		return nil, err
	}
	return navigation, nil

}

func rowsProcessing(rows pgx.Rows) ([]*NavigationItem, error) {
	var result []*NavigationItem

	for rows.Next() {
		var item = &NavigationItem{}
		err := rows.Scan(&item.Id, &item.Item_type, &item.Link, &item.Label, &item.EditTime)
		if err != nil {
			log.Print(err)
			continue
		}

		result = append(result, item)

	}

	return result, nil
}

func (n *NavigationModel) GetAll() ([]*NavigationItem, error) {

	sql := "SELECT * FROM remonttiv2.navigation"

	rows, err := n.DB.Query(context.Background(), sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return rowsProcessing(rows)

}

func (n *NavigationModel) GetAllForUser(userId int) ([]*NavigationItem, error) {

	sql := `SELECT 
				remonttiv2.navigation.id, remonttiv2.navigation.item_type, remonttiv2.navigation.link, 
				remonttiv2.navigation.label, remonttiv2.navigation.edit_time
			FROM 
				remonttiv2.navigation, remonttiv2.permissions 
			WHERE 
				remonttiv2.navigation.id = remonttiv2.permissions.component_id AND
				(remonttiv2.permissions.actions & $1) = $1 AND
				remonttiv2.permissions.user_id = $2;`

	rows, err := n.DB.Query(context.Background(), sql, permissionservice.Actions.VIEW, userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return rowsProcessing(rows)

}

func (n *NavigationModel) GetByType(itemType int) ([]*NavigationItem, error) {

	sql := "SELECT * FROM remonttiv2.navigation WHERE item_type=$1"

	rows, err := n.DB.Query(context.Background(), sql, itemType)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return rowsProcessing(rows)

}

func (n *NavigationModel) GetById(id int) (*NavigationItem, error) {

	sql := "SELECT * FROM remonttiv2.navigation WHERE id=$1"

	row := n.DB.QueryRow(context.Background(), sql, id)

	return rowProcessing(row)

}

// Get Получает строку из таблица navigation по типу пункта меню, сслыке и заголовку
func (n *NavigationModel) Get(itemType int, link, label string) (*NavigationItem, error) {

	sql := "SELECT * FROM remonttiv2.navigation WHERE item_type=$1 AND link=$2 AND label=$3"

	row := n.DB.QueryRow(context.Background(), sql, itemType, link, label)

	return rowProcessing(row)
}

// Create Есть лишние действия -- проверка на существование,
func (n *NavigationModel) Create(itemType int, link, label string) (*NavigationItem, error) {

	item, err := n.Get(itemType, link, label)

	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}
	if item != nil {
		return item, ErrItemAlreadyExists
	}

	editTime := time.Now().Unix()

	sql := `
		INSERT INTO remonttiv2.navigation (item_type, link, label, edit_time)
		VALUES ($1, $2, $3, $4)`

	_, err = n.DB.Exec(context.Background(), sql, itemType, link, label, editTime)
	if err != nil {
		return nil, err
	}

	item, err = n.Get(itemType, link, label)

	if err != nil {
		return nil, err
	}

	return item, nil

}

func (n *NavigationModel) CreateBatch(items []*NavigationItemInsert) error {

	var batch *pgx.Batch

	sql := `
		INSERT INTO remonttiv2.navigation (item_type, link, label, edit_time)
		VALUES ($1, $2, $3, $4)`

	for _, item := range items {
		batch.Queue(sql, item.ItemType, item.Link, item.Label, time.Now().Unix())
	}

	result := n.DB.SendBatch(context.Background(), batch)
	defer result.Close()

	var reserr error
	var res pgconn.CommandTag

	for reserr == nil {
		res, reserr = result.Exec()
		println(res)
	}

	return nil

}
