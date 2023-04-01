package handlers

import (
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

func home(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello! Request was processed")
}
