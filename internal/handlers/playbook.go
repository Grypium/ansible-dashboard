package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/your_username/your_project_name/internal/models"
)

func (app *App) GetPlaybook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	playbook, err := app.DB.GetPlaybook(id)
	if err != nil {
		if err == models.ErrNoPlaybook {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(playbook)
}

func (app *App) CreatePlaybook(w http.ResponseWriter, r *http.Request) {
	var playbook models.Playbook
	err := json.NewDecoder(r.Body).Decode(&playbook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = app.DB.CreatePlaybook(&playbook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(playbook)
}

func (app *App) UpdatePlaybook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var playbook models.Playbook
	err = json.NewDecoder(r.Body).Decode(&playbook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	playbook.ID = id
	err = app.DB.UpdatePlaybook(&playbook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(playbook)
}

func (app *App) DeletePlaybook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = app.DB.DeletePlaybook(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
	