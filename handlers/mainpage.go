package handlers

import (
	"log"
	"net/http"
	"text/template"
)

func mainPage(w http.ResponseWriter, r *http.Request) {

	var uiPath = cfg.ROOT_PATH + "/web/ui/"

	files := []string{
		uiPath + "main.page.gohtml",
		uiPath + "base.layout.gohtml",
		uiPath + "footer.partial.gohtml",
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
