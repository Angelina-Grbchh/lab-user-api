package main

import (
	"log"
	"net/http"

	"github.com/Angelina-Grbchh/lab-user-api/internal/handler"
	"github.com/Angelina-Grbchh/lab-user-api/store"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	st := store.NewUserStore()
	srv := handler.NewServer(st)

	r.HandleFunc("/users", srv.CreateUser).Methods("POST")
	r.HandleFunc("/users", srv.ListUsers).Methods("GET")
	r.HandleFunc("/users/{id}", srv.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", srv.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", srv.DeleteUser).Methods("DELETE")

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", r)
}
