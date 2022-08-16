package database

type Users struct {
	Id                 int
	CompanyId          int
	UserName           string
	Password           string
	LastLoginDate      string
	LastLoginErrorDate string
}
