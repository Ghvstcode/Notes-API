package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	 "github.com/GhvstCode/Notes-API/controllers"
)

func handleRequest(){
	r := mux.NewRouter()

	//Create a new User
	r.HandleFunc("/api/user/new", controllers.NewUser).Methods("POST")
	////Login a User
	//r.HandleFunc("/api/user/login", createNewArticle).Methods("POST")
	////Update a User's info, user provides ID of note they wish to have updated
	//r.HandleFunc("/article/{id}", updateArticle).Methods("PUT")
	////Delete a note
	//r.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
   //
   //
	////Return all a users article
	//r.HandleFunc("/articles", returnAllArticles)
   ////Create a new article
	//r.HandleFunc("/newarticle", createNewArticle).Methods("POST")
	////Update a note, user provides ID of note they wish to have updated
	//r.HandleFunc("/article/{id}", updateArticle).Methods("PUT")
	////Delete a note
	//r.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	////Search for a note by its ID.
	//r.HandleFunc("/article/{id}", returnSingleArticle)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func main(){
	handleRequest()
}
