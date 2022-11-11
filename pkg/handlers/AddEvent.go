package handlers

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"

	"github.com/sebavidal10/api-go-crud/pkg/mocks"
	"github.com/sebavidal10/api-go-crud/pkg/models"
)

func AddEvent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var event models.Event

	json.Unmarshal(body, &event)

	event.Id = rand.Intn(100)
	mocks.Events = append(mocks.Events, event)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Created")
}
