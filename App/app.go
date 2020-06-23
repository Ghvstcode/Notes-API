package App

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"

	"github.com/GhvstCode/Notes-API/Auth"
	"github.com/GhvstCode/Notes-API/controllers"
)

var(
	r = mux.NewRouter()
)

func HandleRequest() {
	r.Use(Auth.JWT)
	//Create a new User
	r.HandleFunc("/api/user/new", controllers.NewUser).Methods(http.MethodPost)
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
}
func Start(){
	l := log.New(os.Stdout, " Notes-API", log.LstdFlags)

	HandleRequest()

	s := &http.Server{
		Addr: ":8080",
		Handler: r,
		//IdleTimeout: 120*time.Second,
		//ReadTimeout: 1*time.Second,
		//WriteTimeout: 1*time.Second,
	}

	go func () {
		err := 	s.ListenAndServe()

		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <- sigChan
	l.Print("Received terminate, graceful shutdown! Signal: ",sig)

	tc, _:= context.WithTimeout(context.Background(), 30*time.Second)
	_ = s.Shutdown(tc)
}
