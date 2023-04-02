package userservice

import (
	"context"
	"errors"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net/http"

	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id                 int
	CompanyId          int
	UserName           string
	Password           string
	LastLoginDate      int
	LastLoginErrorDate int
}

type UserModel struct {
	DB          *pgxpool.Pool
	CookieStore *sessions.CookieStore
}

// ErrUserAlreadyExists Ошибка ErrUserAlreadyExists возвращается при попытке создать
// пользователя который уже существует в системе
var ErrUserAlreadyExists = errors.New("users: User already exists")

func (u *UserModel) GetCurrentUserId(r *http.Request, secret string) (int, error) {
	session, err := u.CookieStore.Get(r, secret)
	if err != nil {
		return 0, err
	}
	return session.Values["userid"].(int), nil
}

// Возвращает пользователя по его userName и companyId
// если пользователь не существует возвращает ошибку ErrNoRows
func (u *UserModel) GetByNameAndCompanyId(userName string, companyId int) (*User, error) {
	sql := `SELECT * 
			FROM 
				remonttiv2.users 
			WHERE 
				user_name=$1 AND company_id=$2;`
	row := u.DB.QueryRow(context.Background(), sql, userName, companyId)

	var user = &User{}
	err := row.Scan(&user.Id, &user.CompanyId, &user.UserName, &user.Password, &user.LastLoginDate, &user.LastLoginErrorDate)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Создает нового пользователя с данными
// userName, password и companyId.
// Если пользователь существует, то
// возвращает ошибку ErrUserAlreadyExists и пользователя
func (u *UserModel) Create(userName string, password string, companyId int) (*User, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return nil, err
	}

	// Пытаемся получить пользователя с таким именем и companyId
	user, err := u.GetByNameAndCompanyId(userName, companyId)

	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}
	if user != nil {
		return user, ErrUserAlreadyExists
	}

	sql := `INSERT INTO remonttiv2.users 
				(company_id, user_name, password, last_login_date, last_login_error_date) 
			VALUES 
				($1, $2, $3, $4, $5);`

	_, err = u.DB.Exec(context.Background(), sql, companyId, userName, string(bytes), 0, 0)
	if err != nil {
		return nil, err
	}

	user, err = u.GetByNameAndCompanyId(userName, companyId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserModel) GetAll() ([]*User, error) {

	sql := `SELECT * FROM remonttiv2.users;`

	rows, err := u.DB.Query(context.Background(), sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return rowsProcessing(rows)
}

func rowsProcessing(rows pgx.Rows) ([]*User, error) {
	var result []*User

	for rows.Next() {
		var item = &User{}
		err := rows.Scan(&item.Id, &item.CompanyId, &item.UserName, &item.Password, &item.LastLoginDate, &item.LastLoginErrorDate)
		if err != nil {
			log.Print(err)
			continue
		}

		result = append(result, item)

	}

	return result, nil
}
