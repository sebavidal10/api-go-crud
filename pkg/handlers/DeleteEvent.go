package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sebavidal10/api-go-crud/pkg/mocks"
)

func DeleteEvent(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for index, event := range mocks.Events {
		if event.Id == id {
			mocks.Events = append(mocks.Events[:index], mocks.Events[index+1:]...)
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode("Deleted")
			return
		}
	}
}
