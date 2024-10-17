package rest

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func Handle(handler func(r *http.Request) Response) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		response := handler(r)
		if response.Error != nil {
			err := response.Error
			fmt.Print(err)
			//TODO: REFACTOR LOG ERROR
		}

		if err := sendJSON(w, response.Body, response.Status); err != nil {
			fmt.Print(err)
			//TODO: REFACTOR LOG ERROR

		}
	}
}

func sendJSON(w http.ResponseWriter, payload any, statusCode int) error {
	if payload == nil {
		w.WriteHeader(statusCode)
		return nil
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(payload)
}

func URLParam(r *http.Request, key string) string {
	return chi.URLParam(r, key)
}
