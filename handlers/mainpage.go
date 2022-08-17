package handlers

import (
	"log"
	"net/http"
	"text/template"

	"github.com/dmidokov/remontti-v2/translations"
)

// Функция обработчик для главной страницы
func mainPage(w http.ResponseWriter, r *http.Request) {

	var rootPath = cfg.ROOT_PATH + "/web/ui/"

	// Список файлов для шаблона
	files := []string{
		rootPath + "main.page.gohtml",
		rootPath + "base.layout.gohtml",
		rootPath + "footers/footer.partial.gohtml",
		rootPath + "headers/mainpage.partial.gohtml",
		rootPath + "bodies/mainpage.partial.gohtml",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	pageData.Title = "Меню"
	pageData.Translation, err = translations.GetTranslations("mainpage", cfg)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, pageData)
	if err != nil {
		log.Print(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

}
