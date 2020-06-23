package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/GhvstCode/Notes-API/Auth"
	"github.com/GhvstCode/Notes-API/controllers"
)

func handleRequest(){
	r := mux.NewRouter()
	r.Use(Auth.JWT)
	//Create a new User
	r.HandleFunc("/api/user/new", controllers.NewUser).Methods("POST")
	//Login a User
	r.HandleFunc("/api/user/login", controllers.LoginUser).Methods("POST")
	//Update a User's info,
	r.HandleFunc("/api/user/update", controllers.UpdateUser).Methods("PUT")
	//Create a new Note
	r.HandleFunc("/api/notes/new", controllers.NewNote).Methods("POST")
	//Return all a users Notes
	r.HandleFunc("/api/notes/GetNotes", controllers.GetAllNotes).Methods("GET")
	//Update a note, user provides ID of note they wish to have updated
	r.HandleFunc("/api/notes/{id}", controllers.UpdateNote).Methods("PUT")
	//Delete a note
	r.HandleFunc("/api/notes/{id}", controllers.DeleteNote).Methods("DELETE")


	log.Fatal(http.ListenAndServe(":8080", r))
}

func main(){
	handleRequest()
}
