package api

import (
	"fmt"
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	_, err := w.Write([]byte(`{"message": "Not Found"}`))
	if err != nil {
		fmt.Errorf("failed to generate response, err=%v", err.Error())
	}
}
