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
		u.RespondJson(w, u.Message(false, "Invalid request", http.StatusBadRequest))
		return
	}
	//-create account by calling create method on the model and send back to the client
	res:= user.Create()
	u.RespondJson(w, res)
}

func LoginUser(w http.ResponseWriter, r *http.Request){
	user := &models.UserAccount{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.RespondJson(w, u.Message(false, "Invalid request", http.StatusBadRequest))
		return
	}
    res := models.Login(user.Email, user.Password)
    u.RespondJson(w, res)
}

func UpdateUser(w http.ResponseWriter, r *http.Request){
	userID := r.Context().Value("user").(uint)
	user :=&models.UserAccount{}

	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.RespondJson(w, u.Message(false, "Invalid request", http.StatusBadRequest))
		return
	}

	updatedUser := user.Update(userID)
	u.RespondJson(w, updatedUser)
	//We search for the ID of that user,
	//We update provided fields &  store it in the DB
	//If provided field is the password, we would want to hash it before saving!
}