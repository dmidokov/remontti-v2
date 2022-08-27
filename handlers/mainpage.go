package handlers

import (
	"log"
	"net/http"
	"text/template"

	"github.com/dmidokov/remontti-v2/navigationservice"
	"github.com/dmidokov/remontti-v2/translationservice"
)

// Функция обработчик для главной страницы
func (hm *HandlersModel) mainPage(w http.ResponseWriter, r *http.Request) {

	var navigation navigationservice.NavigationModel = navigationservice.NavigationModel{DB: hm.DB}
	var translation translationservice.TranslationsModel = translationservice.TranslationsModel{DB: hm.DB}

	var rootPath = hm.Config.ROOT_PATH + "/web/ui/"

	// Список файлов для шаблона
	files := []string{
		rootPath + "main.page.gohtml",
		rootPath + "base.layout.gohtml",
		rootPath + "footers/footer.partial.gohtml",
		rootPath + "headers/mainpage.partial.gohtml",
		rootPath + "bodies/mainpage.partial.gohtml",
		rootPath + "navigations/topnavigation.partial.gohtml",
	}

	var pageData = loginPageData{
		Title:       "",
		Translation: make(map[string]string),
		Navigation:  make(map[string]navigationData),
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	pageData.Exam = "some string"
	pageData.Title = "Меню"
	translations, err := translation.Get("mainpage", hm.Config)

	for _, translation := range translations {
		pageData.Translation[translation.Label] = translation.Ru
	}

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	items, err := navigation.GetAll()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	labels := make(map[string]navigationData)

	for _, item := range items {
		labels[item.Label] = navigationData{
			Link:        item.Link,
			Translation: pageData.Translation[item.Label],
		}
	}

	pageData.Navigation = labels

	err = ts.Execute(w, pageData)
	if err != nil {
		log.Print(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

}
