package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/Angelina-Grbchh/lab-user-api/model"
	"github.com/Angelina-Grbchh/lab-user-api/store"

	"github.com/gorilla/mux"
)

type Server struct {
	Store *store.UserStore
}

func NewServer(s *store.UserStore) *Server {
	return &Server{Store: s}
}

func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)

	// Валідація
	if strings.TrimSpace(user.Name) == "" || len(user.Name) > 255 {
		http.Error(w, "Invalid name", http.StatusBadRequest)
		return
	}

	if user.Age < 18 {
		http.Error(w, "User must be at least 18 years old", http.StatusBadRequest)
		return
	}

	created, err := s.Store.AddUser(r.Context(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

func (s *Server) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, _ := s.Store.ListUsers(r.Context())
	json.NewEncoder(w).Encode(map[string][]model.User{"users": users})
}

func (s *Server) GetUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	user, err := s.Store.GetUser(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (s *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)

	if strings.TrimSpace(user.Name) == "" || len(user.Name) > 255 {
		http.Error(w, "Invalid name", http.StatusBadRequest)
		return
	}

	if user.Age < 18 {
		http.Error(w, "User must be at least 18 years old", http.StatusBadRequest)
		return
	}

	updated, err := s.Store.UpdateUser(r.Context(), id, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(updated)
}

func (s *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	err := s.Store.DeleteUser(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
