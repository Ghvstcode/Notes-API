package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/GhvstCode/Notes-API/models"
	u "github.com/GhvstCode/Notes-API/utils"
)

//Create new user handler function!
func NewUser(w http.ResponseWriter, r *http.Request){
	//-new instance of the new user model
	user := &models.UserAccount{}
	//-decode the new provided details
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		//fill the message util and send back the error!
		u.RespondJson(w, u.Message(false, "Invalid request"))
		return
	}
	//-create account by calling create method on the model and send back to the client
	res, _ := user.Create()
	u.RespondJson(w, res)
}
