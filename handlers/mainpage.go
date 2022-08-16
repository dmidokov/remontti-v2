package handlers

import (
	"log"
	"net/http"
	"text/template"
)

// Функция обработчик для главной страницы
func mainPage(w http.ResponseWriter, r *http.Request) {

	var rootPath = cfg.ROOT_PATH + "/web/ui/"

	// Список файлов для шаблона
	files := []string{
		rootPath + "main.page.gohtml",
		rootPath + "base.layout.gohtml",
		rootPath + "footer.partial.gohtml",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

}
