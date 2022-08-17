package handlers

import (
	"log"
	"net/http"
)

func auth(f http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		session, err := sessionStore.Get(r, "session-key")
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
		}

		if auth, ok := session.Values["authenticated"].(bool); ok && auth {
			session.Options.MaxAge = 3600
			session.Save(r, w)
			f(w, r)
		} else {
			login(w, r)
		}

	}

}