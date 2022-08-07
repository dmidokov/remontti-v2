package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Функция не имеет в основе функционала
// просто проверка 
func registration(router *mux.Router) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type RegistrationInfo struct {
			ModuleName string
			ModuleLink string
			Token      string
		}

		var regInfo RegistrationInfo

		var decoder = json.NewDecoder(r.Body)

		var _ = decoder.Decode(&regInfo)

		router.HandleFunc(regInfo.ModuleLink, getRequest)

		fmt.Fprint(w, regInfo)
	}
}

func getRequest(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello! Request was processed for new module link")
}
