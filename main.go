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

//"Authorization: Bearer {ACCESS_TOKEN}"
//curl -X PUT -H "Content-Type: application/json" "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjJ9.6KwkH6ycpEbtdfZuM1N_m8E_W1UhQ5cV_SMRrqvAO40" -d '{"name":"aupdated","email":"aupdated@gmail.com"}' http://localhost:8080/api/user/update