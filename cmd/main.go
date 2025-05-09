package main

import (
    "lab-user-api/internal/handler"
    "lab-user-api/store"
    "log"
    "net/http"

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
