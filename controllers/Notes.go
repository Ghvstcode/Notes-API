package controllers

import (
	"encoding/json"
	"net/http"

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
