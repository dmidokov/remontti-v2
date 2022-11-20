package handlers

import (
	"log"
	"net/http"
)

func (hm *HandlersModel) auth2(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := hm.CookieStore.Get(r, hm.Config.SESSIONS_SECRET)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}

		if auth, ok := session.Values["authenticated"].(bool); ok && auth {
			session.Options.MaxAge = 3600
			session.Save(r, w)
			f(w, r)
		} else {
			hm.login(w, r)
		}
	}
}

func (hm *HandlersModel) auth(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := hm.CookieStore.Get(r, hm.Config.SESSIONS_SECRET)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}

		if auth, ok := session.Values["authenticated"].(bool); ok && auth {
			session.Options.MaxAge = 3600
			session.Save(r, w)
			// Не кешируем поход на главную, иначе FF редиректит на разлогине на главную
			// и пытается показать ее из кеша
			// надо прочекать как работает в ФФ
			if r.URL.String() == "/" {
				w.Header().Set("cache-control", "no-cache")
			}
			f.ServeHTTP(w, r)
		} else {
			http.Redirect(w, r, "/login/", http.StatusSeeOther)
			//handleFileServer(hm.Config.ROOT_PATH+"/web/vueui/remontti-ui/dist/login/", r.RequestURI).ServeHTTP(w, r)
		}
	}
}
