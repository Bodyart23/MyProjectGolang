package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"somePriject/web-app/controller"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", controller.GetUsers).Methods("GET")
	router.HandleFunc("/users/filter/", controller.UsersFilter).Methods("GET")
	router.HandleFunc("/user/", controller.CreateUser).Methods("POST")
	router.HandleFunc("/user/", controller.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/", controller.DeleteUser).Methods("DELETE")

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
