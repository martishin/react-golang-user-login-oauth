package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/martishin/react-golang-goth-auth/internal/database"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func HealthHandler(db database.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jsonResp, _ := json.Marshal(db.Health())
		_, _ = w.Write(jsonResp)
	}
}
