package controllers

import (
	"encoding/json"
	"github.com/GhvstCode/Notes-API/models"
	"io/ioutil"
	"net/http"
)

//Create new user handler function!
func NewUser(w http.ResponseWriter, r *http.Request){
	//-new instance of the new user model
	user := &models.UserAccount{}
	//-decode the new provided details
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//fill the message util and send back the error!
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	_ = json.Unmarshal(reqBody, user)
	err = json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	//-create account and send back to the client
	res := user.Create()
	u.Respond(w, res)
}
