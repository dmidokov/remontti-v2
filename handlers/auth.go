package handlers

import (
	"log"
	"net/http"
)

func (hm *HandlersModel) auth(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := hm.CookieStore.Get(r, hm.Config.SESSIONS_SECRET)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
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

func (hm *HandlersModel) auth1(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := hm.CookieStore.Get(r, hm.Config.SESSIONS_SECRET)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
		}

		if auth, ok := session.Values["authenticated"].(bool); ok && auth {
			session.Options.MaxAge = 3600
			session.Save(r, w)
			f.ServeHTTP(w, r)
		} else {
			println("no auth")
			handleFileServer(hm.Config.ROOT_PATH+"/web/vueui/remontti-ui/dist/login/", r.RequestURI).ServeHTTP(w, r)
		}
	}
}
