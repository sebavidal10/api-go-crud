package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sebavidal10/api-go-crud/pkg/mocks"
	"github.com/sebavidal10/api-go-crud/pkg/models"
)

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedEvent models.Event
	json.Unmarshal(body, &updatedEvent)

	for index, event := range mocks.Events {
		if event.Id == id {
			event.Title = updatedEvent.Title
			event.Description = updatedEvent.Description
			event.Address = updatedEvent.Address

			mocks.Events[index] = event

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode("Updated")
			return
		}
	}
}
