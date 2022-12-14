package database

import (
	"github.com/dmidokov/remontti-v2/companyservice"
	"github.com/dmidokov/remontti-v2/config"
	"github.com/dmidokov/remontti-v2/navigationservice"
	"github.com/dmidokov/remontti-v2/permissionservice"
	"github.com/dmidokov/remontti-v2/translationservice"
	"github.com/dmidokov/remontti-v2/userservice"
)

type usersList []*userservice.User
type navigationList []*navigationservice.NavigationItem
type permissionsList []*permissionservice.Permissons
type translationsList []*translationservice.Translation
type companiesList []*companyservice.Company

func GetCompaniesDataToInsert() companiesList {
	return companiesList{
		&companyservice.Company{CompanyId: 0, CompanyName: "CONTROL", HostName: "control.remontti.site", EditTime: 0},
	}
}

func GetUserDataToInsert(cfg *config.Configuration) usersList {

	return usersList{
		&userservice.User{CompanyId: 0, UserName: "admin", Password: cfg.ADMIN_PASSWORD, LastLoginDate: 0, LastLoginErrorDate: 0},
	}

}

func GetNavigationDataToInsert(cfg *config.Configuration) navigationList {

	return navigationList{
		&navigationservice.NavigationItem{Item_type: 1, Link: "/", Label: "Home", EditTime: 0},
		&navigationservice.NavigationItem{Item_type: 1, Link: "/companies/", Label: "Companies", EditTime: 0},
		&navigationservice.NavigationItem{Item_type: 1, Link: "/settings/", Label: "Settings", EditTime: 0},
		&navigationservice.NavigationItem{Item_type: 1, Link: "/logout", Label: "Logout", EditTime: 0},
	}

}

func (pg *DatabaseModel) GetPermissionsDataToInsert(cfg *config.Configuration) (permissionsList, error) {
	var userService = userservice.UserModel{DB: pg.DB}
	var navigationService = navigationservice.NavigationModel{DB: pg.DB}
	var companiesService = companyservice.CompanyModel{DB: pg.DB}

	user, err := userService.GetByNameAndCompanyId("admin", 0)
	if err != nil {
		return nil, err
	}

	var permissionList = permissionsList{}
	var actions = permissionservice.Actions

	navigationItems, err := navigationService.GetAll()
	if err != nil {
		return nil, err
	}

	for _, item := range navigationItems {
		permissionList = append(permissionList, &permissionservice.Permissons{ComponentId: item.Id, ComponentType: "navigation", UserId: user.Id, Actions: actions.VIEW | actions.EDIT | actions.DELETE})
	}

	companiesItems, err := companiesService.GetAll()
	if err != nil {
		return nil, err
	}

	for _, item := range companiesItems {
		permissionList = append(permissionList, &permissionservice.Permissons{ComponentId: item.CompanyId, ComponentType: "company", UserId: user.Id, Actions: actions.VIEW | actions.EDIT | actions.DELETE})
	}

	return permissionList, nil
}

func GetTranslationsDataToInsert(cfg *config.Configuration) translationsList {

	result := []*translationservice.Translation{}

	var loginPage = translationsList{
		&translationservice.Translation{Name: "loginpage", Label: "LoginFieldHeader", Ru: "??????????:", En: "", EditTime: 0},
		&translationservice.Translation{Name: "loginpage", Label: "PasswordFieldHeader", Ru: "????????????:", En: "", EditTime: 0},
		&translationservice.Translation{Name: "loginpage", Label: "SignIn", Ru: "??????????", En: "", EditTime: 0},
		&translationservice.Translation{Name: "loginpage", Label: "LoginFormHeader", Ru: "????????", En: "", EditTime: 0},
		&translationservice.Translation{Name: "loginpage", Label: "PageTitle", Ru: "????????", En: "", EditTime: 0},
		&translationservice.Translation{Name: "loginpage", Label: "InvalidData", Ru: "?????????????????????? ????????????", En: "", EditTime: 0},
		&translationservice.Translation{Name: "loginpage", Label: "UserIsNotExists", Ru: "???????????????????????? ?? ?????????? ???????????? ???? ????????????????????", En: "", EditTime: 0},
		&translationservice.Translation{Name: "loginpage", Label: "ErrorTryAgain", Ru: "????????????, ???????????????????? ?????? ??????", En: "", EditTime: 0},
		&translationservice.Translation{Name: "loginpage", Label: "InvalidUserOrPassword", Ru: "???????????????? ?????????? ?????? ????????????", En: "", EditTime: 0},
		&translationservice.Translation{Name: "loginpage", Label: "EmptyLoginOrPassword", Ru: "?????????? ?????? ???????????? ???? ??????????????", En: "", EditTime: 0},
		&translationservice.Translation{Name: "loginpage", Label: "LoginIsEmpty", Ru: "???? ???????????? ??????????", En: "", EditTime: 0},
		&translationservice.Translation{Name: "loginpage", Label: "PasswordIsEmpty", Ru: "???? ???????????? ????????????", En: "", EditTime: 0},
	}

	var mainPage = translationsList{
		&translationservice.Translation{Name: "mainpage", Label: "Title", Ru: "??????????????", En: "", EditTime: 0},
		&translationservice.Translation{Name: "mainpage", Label: "Account", Ru: "???????????? ??????????????", En: "", EditTime: 0},
		&translationservice.Translation{Name: "mainpage", Label: "Settings", Ru: "??????????????????", En: "", EditTime: 0},
		&translationservice.Translation{Name: "mainpage", Label: "Logout", Ru: "??????????", En: "", EditTime: 0},
	}

	var navigationPage = translationsList{
		&translationservice.Translation{Name: "navigation", Label: "LoginFieldHeader", Ru: "????????...", En: "", EditTime: 0},
		&translationservice.Translation{Name: "navigation", Label: "AccountSettings", Ru: "???????????????????? ??????????????????", En: "", EditTime: 0},
		&translationservice.Translation{Name: "navigation", Label: "Logout", Ru: "??????????", En: "", EditTime: 0},
		&translationservice.Translation{Name: "navigation", Label: "Settings", Ru: "??????????????????", En: "", EditTime: 0},
		&translationservice.Translation{Name: "navigation", Label: "Home", Ru: "??????????????", En: "", EditTime: 0},
		&translationservice.Translation{Name: "navigation", Label: "Companies", Ru: "??????????????????????", En: "", EditTime: 0},
		&translationservice.Translation{Name: "navigation", Label: "Account", Ru: "???????????? ??????????????", En: "", EditTime: 0},
	}

	var companiesPage = translationsList{
		&translationservice.Translation{Name: "companies", Label: "EditCompany", Ru: "??????????????????????????", En: ""},
		&translationservice.Translation{Name: "companies", Label: "DeleteCompany", Ru: "??????????????", En: ""},
		&translationservice.Translation{Name: "companies", Label: "CompanyTableTitle", Ru: "????????????????", En: ""},
		&translationservice.Translation{Name: "companies", Label: "AddCompany", Ru: "????????????????", En: ""},
		&translationservice.Translation{Name: "companies", Label: "CompanyName", Ru: "????????????????", En: ""},
		&translationservice.Translation{Name: "companies", Label: "CompanyHost", Ru: "????????", En: ""},
	}

	result = append(result, loginPage...)
	result = append(result, mainPage...)
	result = append(result, navigationPage...)
	result = append(result, companiesPage...)

	return result
}
