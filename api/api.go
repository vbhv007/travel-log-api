package api

import (
	"encoding/json"
	"net/http"
)

func RootPage(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello World!")
}
