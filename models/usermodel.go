package models

import (
	"fmt"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"

	u "github.com/GhvstCode/Notes-API/utils"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}

//a struct to rep user account
type UserAccount struct {
	gorm.Model
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token";sql:"-"`
}
//Function for creating a new token!
func genAuthToken(ua *UserAccount)(string, error){
	t := &Token{UserId: ua.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), t)
	tokenString, err := token.SignedString([]byte(os.Getenv("token_password")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
//Function for validating a new user before saving
func (ua *UserAccount) Validate() (map[string]interface{}, bool) {

//Check to see if Email is valid
	if !strings.Contains(ua.Email, "@") {
		return u.Message(false, "Please provide a valid email address"), false
	}
//Check to see if password is valid!
	if len(ua.Password) < 6 {
		return u.Message(false, "Password is required"), false
	}
//Check to see if password is secure
	if strings.Contains(ua.Password, "abcdefg") {
		return u.Message(false, "Please provide a valid password"), false
	}
//Check to see if Email is unique
	temp := &UserAccount{}

	//check for errors and duplicate emails
	err := GetDB().Table("user_accounts").Where("email = ?", ua.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		//log.Fatal(err)
		fmt.Print("Hello", err)
		return u.Message(false, "Connection error. Please retry"), false
	}
	if temp.Email != "" {
		return u.Message(false, "Email address already in use by another user."), false
	}

	return u.Message(false, "Validated"), true
}

//Function for creating the new User
func (ua *UserAccount) Create() (map[string]interface{}, bool){
//Check if the values provided are valid
	if resp, b := ua.Validate(); !b {
		return resp, false
	}

//Hash the password!
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(ua.Password), bcrypt.DefaultCost)
	if err != nil {
		return u.Message(false, "An error occurred! Unable to save user"), false
	}

	ua.Password = string(hashedPassword)
	GetDB().Create(ua)

	if ua.ID <= 0 {
		fmt.Print("Na me")
		return u.Message(false, "Failed to create account, connection error."), false
	}

	 t, e := genAuthToken(ua)
	if e != nil {
		return u.Message(false, "Failed to create account, connection error."), false
	}

	ua.Token = t
	ua.Password = ""

	response := u.Message(true, "Account has been created")
	response["account"] = ua
	return response, true
}