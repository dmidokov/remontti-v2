package handlers

import (
	"encoding/json"
	"github.com/dmidokov/remontti-v2/translationservice"
	"net/http"
	"strings"
)

func (hm *HandlersModel) getTranslationsApi(w http.ResponseWriter, r *http.Request) {
	if hm.Config.MODE == "dev" {
		if r.Method == "OPTIONS" {
			setCorsHeaders(&w, r)
			return
		}
		setCorsHeaders(&w, r)
	}

	pages := r.URL.Query().Get("pages")

	if pages == "" {
		json.NewEncoder(w).Encode(response{
			Status:  "error",
			Message: "pageListIsEmpty",
			Errors:  []string{"pages list is empty"}})
	}

	var translation = translationservice.TranslationsModel{DB: hm.DB}
	var translations []*translationservice.Translation
	var cachedResult = translation.Pop(pages)
	var err error

	if cachedResult != nil {
		translations = cachedResult.Result
	} else {
		sliceOfPages := strings.Split(pages, ",")

		translations, err = translation.Get(sliceOfPages...)

		translation.Push(pages, translations)

		if err != nil {
			json.NewEncoder(w).Encode(response{
				Status:  "error",
				Message: "cantGetTranslations",
				Errors:  []string{"can't get translations for the pages"}})
		}
	}

	var result = make(map[string]string)

	for _, item := range translations {
		result[item.Label] = item.Ru
	}

	json.NewEncoder(w).Encode(result)
}
