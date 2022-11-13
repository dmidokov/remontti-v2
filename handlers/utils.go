package handlers

import (
	"errors"
	"net/http"

	"github.com/dmidokov/remontti-v2/companyservice"
	"github.com/dmidokov/remontti-v2/navigationservice"
	"github.com/dmidokov/remontti-v2/translationservice"
)

type SessionData struct {
	UserId    int
	CompanyId int
}

func (hm *HandlersModel) getUserNavigation(userId int, translation map[string]string) (map[string]navigationData, error) {

	var navigation navigationservice.NavigationModel = navigationservice.NavigationModel{DB: hm.DB}

	items, err := navigation.GetAllForUser(userId)
	if err != nil {
		return nil, err
	}

	labels := make(map[string]navigationData)

	for _, item := range items {
		labels[item.Label] = navigationData{
			Link:        item.Link,
			Translation: translation[item.Label],
		}
	}

	return labels, nil
}

func (hm *HandlersModel) getSessionData(r *http.Request) (*SessionData, error) {
	session, _ := hm.CookieStore.Get(r, hm.Config.SESSIONS_SECRET)

	userId, exist := session.Values["userid"].(int)
	if !exist {

		return nil, errors.New("value is not exists")
	}

	companyId, exist := session.Values["companyid"].(int)
	if !exist {
		return nil, errors.New("value is not exists")
	}

	return &SessionData{UserId: userId, CompanyId: companyId}, nil

}

func (hm *HandlersModel) getTranslations(pagenames ...string) (map[string]string, error) {

	var translation translationservice.TranslationsModel = translationservice.TranslationsModel{DB: hm.DB}

	translations, err := translation.Get(pagenames...)

	if err != nil {
		return nil, err
	}

	var result = make(map[string]string)

	for _, translation := range translations {
		result[translation.Label] = translation.Ru
	}

	return result, nil
}

func (hm *HandlersModel) getCompanies() (*[]companiesData, error) {

	var companyservice companyservice.CompanyModel = companyservice.CompanyModel{DB: hm.DB}

	companies, err := companyservice.GetAll()
	if err != nil {
		return nil, err
	}

	var result = make([]companiesData, 0)

	for _, company := range companies {
		result = append(result, companiesData{Name: company.CompanyName, Host: company.HostName, ID: company.CompanyId})
	}

	return &result, nil
}
