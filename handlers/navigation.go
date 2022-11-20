package handlers

import (
	"encoding/json"
	"github.com/dmidokov/remontti-v2/navigationservice"
	"github.com/dmidokov/remontti-v2/userservice"
	"net/http"
)

type NavigationResult struct {
	Link  string `json:"link"`
	Label string `json:"label"`
}

func (hm *HandlersModel) getNavigationApi(w http.ResponseWriter, r *http.Request) {
	var log = hm.Logger

	if hm.Config.MODE == "dev" {
		if r.Method == "OPTIONS" {
			setCorsHeaders(&w, r)
			return
		}
		setCorsHeaders(&w, r)
	}

	var navigationService = navigationservice.NavigationModel{DB: hm.DB}
	var userService = userservice.UserModel{DB: hm.DB, CookieStore: hm.CookieStore}

	userId, err := userService.GetCurrentUserId(r, hm.Config.SESSIONS_SECRET)
	if err != nil {
		// TODO: Вынести подобного рода обработчики в отдельный пакет или хотя бы файл
		err := json.NewEncoder(w).Encode(response{
			Status:  "error",
			Message: pageData.Translation["ErrorTryAgain"],
			Errors:  []string{"Internal server error"}})

		if err != nil {
			log.Error("Не удалось кодировать JSON: %s", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}

	navigationItems, err := navigationService.GetAllForUser(userId)
	if err != nil {
		err := json.NewEncoder(w).Encode(response{
			Status:  "error",
			Message: pageData.Translation["ErrorTryAgain"],
			Errors:  []string{"Internal server error"}})

		if err != nil {
			log.Error("Не удалось кодировать JSON: %s", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}

	result := make([]*NavigationResult, 0)
	for _, item := range navigationItems {
		result = append(
			result,
			&NavigationResult{Link: item.Link, Label: item.Label},
		)
	}

	json.NewEncoder(w).Encode(result)

}
