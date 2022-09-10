package handlers

import (
	"log"
	"net/http"
	"text/template"
)

type mainPageData struct {
	Title       string
	Translation map[string]string
	Navigation  map[string]navigationData
}

// Функция обработчик для главной страницы
func (hm *HandlersModel) mainPage(w http.ResponseWriter, r *http.Request) {

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

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// Структура с данными для страницы
	var pageData = mainPageData{
		Title:       "",
		Translation: make(map[string]string),
		Navigation:  make(map[string]navigationData),
	}

	pageData.Title = "Меню" // нужно ли???

	// Получаем переводы для нужных страниц, в данном случае
	// для главной страницы, а также для верхнего навигационного меню
	pageData.Translation, err = hm.getTranslations("mainpage", "navigation")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

	// получаем данные сессии, а иименно
	// идентификаторы пользователя и компании
	sessionsData, err := hm.getSessionData(r)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

	// Получаем навигационное меню, точнее ссылки и переводы
	// На вход подаются также и переводы, так как навигация в таблицах
	// хранится как ссылки и лэйблы, и для верного отображения необходимо
	// сматчить лейблы навигации и лейблы переводов
	pageData.Navigation, err = hm.getUserNavigation(sessionsData.UserId, pageData.Translation)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

	// Выполняем шаблон, стоит не забыть закешировать шаблоны в будущем
	err = ts.Execute(w, pageData)
	if err != nil {
		log.Print(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

}
