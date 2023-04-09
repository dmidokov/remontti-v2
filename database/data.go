package database

import (
	"github.com/dmidokov/remontti-v2/companyservice"
	"github.com/dmidokov/remontti-v2/config"
	"github.com/dmidokov/remontti-v2/navigationservice"
	"github.com/dmidokov/remontti-v2/permissionservice"
	"github.com/dmidokov/remontti-v2/translationservice"
	"github.com/dmidokov/remontti-v2/userservice"
)

type UsersList []*userservice.User
type NavigationList []*navigationservice.NavigationItem
type PermissionsList []*permissionservice.Permissions
type TranslationsList []*translationservice.Translation
type CompaniesList []*companyservice.Company
type GroupsList []*permissionservice.Group

func GetGroupsDataToInsert() GroupsList {
	return GroupsList{
		&permissionservice.Group{GroupId: 0, GroupName: "Company admin"},
	}
}

func GetCompaniesDataToInsert() CompaniesList {
	return CompaniesList{
		&companyservice.Company{CompanyId: 0, CompanyName: "CONTROL", HostName: "control.remontti.site", EditTime: 0},
	}
}

func GetUserDataToInsert(cfg *config.Configuration) UsersList {

	return UsersList{
		&userservice.User{CompanyId: 0, UserName: "admin", Password: cfg.ADMIN_PASSWORD, LastLoginDate: 0, LastLoginErrorDate: 0},
	}

}

func GetNavigationDataToInsert(cfg *config.Configuration) NavigationList {

	return NavigationList{
		&navigationservice.NavigationItem{Item_type: 1, Link: "/", Label: "Home", EditTime: 0, Ordinal_number: 0},
		&navigationservice.NavigationItem{Item_type: 1, Link: "/companies/", Label: "Companies", EditTime: 0, Ordinal_number: 0},
		&navigationservice.NavigationItem{Item_type: 1, Link: "/settings/", Label: "Settings", EditTime: 0, Ordinal_number: 0},
		&navigationservice.NavigationItem{Item_type: 1, Link: "/logout", Label: "Logout", EditTime: 0, Ordinal_number: 0},
		&navigationservice.NavigationItem{Item_type: 1, Link: "/permissions/", Label: "Permissions", EditTime: 0, Ordinal_number: 0},
	}

}

func (pg *DatabaseModel) GetPermissionsDataToInsert(cfg *config.Configuration) (PermissionsList, error) {
	var userService = userservice.UserModel{DB: pg.DB}
	var navigationService = navigationservice.NavigationModel{DB: pg.DB}
	var companiesService = companyservice.CompanyModel{DB: pg.DB}

	user, err := userService.GetByNameAndCompanyId("admin", 0)
	if err != nil {
		return nil, err
	}

	var permissionList = PermissionsList{}
	var actions = permissionservice.Actions

	navigationItems, err := navigationService.GetAll()
	if err != nil {
		return nil, err
	}

	for _, item := range navigationItems {
		permissionList = append(permissionList, &permissionservice.Permissions{ComponentId: item.Id, ComponentType: "navigation", UserId: user.Id, Actions: actions.VIEW | actions.EDIT | actions.DELETE})
	}

	companiesItems, err := companiesService.GetAll()
	if err != nil {
		return nil, err
	}

	for _, item := range companiesItems {
		permissionList = append(permissionList, &permissionservice.Permissions{ComponentId: item.CompanyId, ComponentType: "company", UserId: user.Id, Actions: actions.VIEW | actions.EDIT | actions.DELETE})
	}

	return permissionList, nil
}

func GetTranslationsDataToInsert(cfg *config.Configuration) TranslationsList {

	var result []*translationservice.Translation

	var loginPage = TranslationsList{
		&translationservice.Translation{Name: "loginpage", Label: "LoginFieldHeader", Ru: "Логин:", En: "", EditTime: 0},
		&translationservice.Translation{Name: "loginpage", Label: "PasswordFieldHeader", Ru: "Пароль:", En: "", EditTime: 0},
		&translationservice.Translation{Name: "loginpage", Label: "SignIn", Ru: "Войти", En: "", EditTime: 0},
		&translationservice.Translation{Name: "loginpage", Label: "LoginFormHeader", Ru: "Вход", En: "", EditTime: 0},
		&translationservice.Translation{Name: "loginpage", Label: "PageTitle", Ru: "Вход", En: "", EditTime: 0},
		&translationservice.Translation{Name: "loginpage", Label: "InvalidData", Ru: "Некоректные данные", En: "", EditTime: 0},
		&translationservice.Translation{Name: "loginpage", Label: "UserIsNotExists", Ru: "Пользователя с таким именем не существует", En: "", EditTime: 0},
		&translationservice.Translation{Name: "loginpage", Label: "ErrorTryAgain", Ru: "Ошибка, попробуйте еще раз", En: "", EditTime: 0},
		&translationservice.Translation{Name: "loginpage", Label: "InvalidUserOrPassword", Ru: "Неверный логин или пароль", En: "", EditTime: 0},
		&translationservice.Translation{Name: "loginpage", Label: "EmptyLoginOrPassword", Ru: "Логин или пароль не указаны", En: "", EditTime: 0},
		&translationservice.Translation{Name: "loginpage", Label: "LoginIsEmpty", Ru: "Не указан логин", En: "", EditTime: 0},
		&translationservice.Translation{Name: "loginpage", Label: "PasswordIsEmpty", Ru: "Не указан пароль", En: "", EditTime: 0},
	}

	var mainPage = TranslationsList{
		&translationservice.Translation{Name: "mainpage", Label: "Title", Ru: "Главная", En: "", EditTime: 0},
		&translationservice.Translation{Name: "mainpage", Label: "Account", Ru: "Личный кабинет", En: "", EditTime: 0},
		&translationservice.Translation{Name: "mainpage", Label: "Settings", Ru: "Настройки", En: "", EditTime: 0},
		&translationservice.Translation{Name: "mainpage", Label: "Logout", Ru: "Выход", En: "", EditTime: 0},
	}

	var navigationPage = TranslationsList{
		&translationservice.Translation{Name: "navigation", Label: "LoginFieldHeader", Ru: "Меню...", En: "", EditTime: 0},
		&translationservice.Translation{Name: "navigation", Label: "AccountSettings", Ru: "Управление аккаунтом", En: "", EditTime: 0},
		&translationservice.Translation{Name: "navigation", Label: "Logout", Ru: "Выход", En: "", EditTime: 0},
		&translationservice.Translation{Name: "navigation", Label: "Settings", Ru: "Настройки", En: "", EditTime: 0},
		&translationservice.Translation{Name: "navigation", Label: "Home", Ru: "Главная", En: "", EditTime: 0},
		&translationservice.Translation{Name: "navigation", Label: "Companies", Ru: "Организации", En: "", EditTime: 0},
		&translationservice.Translation{Name: "navigation", Label: "Account", Ru: "Личный кабинет", En: "", EditTime: 0},
		&translationservice.Translation{Name: "navigation", Label: "CompanyAlreadyExists", Ru: "Компания уже существует", En: "", EditTime: 0},
		&translationservice.Translation{Name: "navigation", Label: "Permissions", Ru: "Настройка прав", En: "", EditTime: 0},
	}

	var companiesPage = TranslationsList{
		&translationservice.Translation{Name: "companies", Label: "EditCompany", Ru: "Редактировать", En: ""},
		&translationservice.Translation{Name: "companies", Label: "DeleteCompany", Ru: "Удалить", En: ""},
		&translationservice.Translation{Name: "companies", Label: "CompanyTableTitle", Ru: "Компании", En: ""},
		&translationservice.Translation{Name: "companies", Label: "Add" +
			"Company", Ru: "Добавить", En: ""},
		&translationservice.Translation{Name: "companies", Label: "CompanyName", Ru: "Название", En: ""},
		&translationservice.Translation{Name: "companies", Label: "CompanyHost", Ru: "Хост", En: ""},
		&translationservice.Translation{Name: "companies", Label: "CompanyAdminName", Ru: "Логин администратора", En: ""},
		&translationservice.Translation{Name: "companies", Label: "CompanyAdminPassword", Ru: "Пароль администратора", En: ""},
	}

	result = append(result, loginPage...)
	result = append(result, mainPage...)
	result = append(result, navigationPage...)
	result = append(result, companiesPage...)

	return result
}
