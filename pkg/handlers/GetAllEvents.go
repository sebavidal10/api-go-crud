package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sebavidal10/api-go-crud/pkg/mocks"
)

func GettAllEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Events)
}
