package handlers

import (
	"encoding/json"
	"github.com/dmidokov/remontti-v2/translationservice"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func (hm *HandlersModel) getTranslationsApi(w http.ResponseWriter, r *http.Request) {

	log := hm.Logger

	if hm.Config.MODE == "dev" {
		if r.Method == "OPTIONS" {
			setCorsHeaders(&w, r)
			return
		}
		setCorsHeaders(&w, r)
	}

	pages := r.URL.Query().Get("pages")

	log.WithFields(logrus.Fields{
		"pages": pages,
	}).Info("Get translations")

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

		if err != nil {
			log.Warning("can't get translations for the pages", err)
			json.NewEncoder(w).Encode(response{
				Status:  "error",
				Message: "cantGetTranslations",
				Errors:  []string{"can't get translations for the pages, ", err.Error()}})
			return
		}

		if (translations) == nil {
			log.Warning("translations list is empty", err)
			json.NewEncoder(w).Encode(response{
				Status:  "error",
				Message: "cantGetTranslations",
				Errors:  []string{"can't get translations for the pages -- empty list"}})
			return
		}

		translation.Push(pages, translations)
	}

	var result = make(map[string]string)

	for _, item := range translations {
		result[item.Label] = item.Ru
	}

	json.NewEncoder(w).Encode(result)
}
