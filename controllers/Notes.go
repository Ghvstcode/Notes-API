package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/GhvstCode/Notes-API/models"
	u "github.com/GhvstCode/Notes-API/utils"
)

func NewNote(w http.ResponseWriter, r *http.Request){
	userID := r.Context().Value("user").(uint)
	//Content goes  here!
	note := &models.Note{}

	err := json.NewDecoder(r.Body).Decode(note)
	if err != nil {
		u.RespondJson(w, u.Message(false, "An Error occurred"))
		return
	}

	note.UserID = userID
	resp := note.Create()
	u.RespondJson(w, resp)
}

func GetAllNotes(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("user").(uint)
	data := models.GetNotes(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.RespondJson(w, resp)
}

func UpdateNote(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]

	note :=&models.Note{}

	err := json.NewDecoder(r.Body).Decode(note)
	if err != nil {
		u.RespondJson(w, u.Message(false, "Invalid request"))
		return
	}

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		u.RespondJson(w, u.Message(false, " An Error Occurred"))
		return
	}
	updatedNote := note.Update(uint(i))
	u.RespondJson(w, updatedNote)
}

func DeleteNote(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]

	note :=&models.Note{}

	err := json.NewDecoder(r.Body).Decode(note)
	if err != nil {
		u.RespondJson(w, u.Message(false, "Invalid request"))
		return
	}

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		u.RespondJson(w, u.Message(false, " An Error Occurred"))
		return
	}
	DeletedNote := note.DeleteNote(uint32(i))
	u.RespondJson(w, DeletedNote)
}