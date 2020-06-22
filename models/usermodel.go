package models

import (
	"fmt"
	"log"
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

//a struct to represent a User
type UserAccount struct {
	gorm.Model
	Name string `json:"name"`
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
//Function for hashing a password
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

//Function for validating a new user before saving
func (ua *UserAccount) Validate() (map[string]interface{}, bool) {
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

	return u.Message(true, "Validated"), true

}

//Function for creating the new User
func (ua *UserAccount) Create() map[string]interface{}{
//Check if the values provided are valid
	if resp, b := ua.Validate(); !b {
		return resp
	}

//Hash the password!
	hashedPassword, err := Hash(ua.Password)
	if err != nil {
		return u.Message(false, "An error occurred! Unable to save user")
	}

	ua.Password = string(hashedPassword)
	GetDB().Create(ua)

	if ua.ID <= 0 {
		fmt.Print("Na me")
		return u.Message(false, "Failed to create account, connection error.")
	}

	 t, e := genAuthToken(ua)
	if e != nil {
		return u.Message(false, "Failed to create account, connection error.")
	}

	ua.Token = t
	ua.Password = ""

	response := u.Message(true, "Account has been created")
	response["account"] = ua
	return response
}

func Login(email string,  password string) map[string]interface{} {
	user := &UserAccount{}
	err := GetDB().Table("user_accounts").Where("email = ?", email).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email address not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return u.Message(false, "Invalid login credentials. Please try again")
	}
	user.Password = ""

	token, _ := genAuthToken(user)
	user.Token = token

	resp := u.Message(true, "Logged In")
	resp["account"] = user
	return resp
}

func (ua *UserAccount)Update(userID uint) map[string]interface{} {
	//Hash the password!
	hashedPassword, err := Hash(ua.Password)
	if err != nil {
		return u.Message(false, "An error occurred! Unable to save user")
	}

	ua.Password = string(hashedPassword)

	db = GetDB().Debug().Model(&UserAccount{}).Where("id = ?", userID).Take(&UserAccount{}).UpdateColumns(
		map[string]interface{}{
			"password":  ua.Password,
			"name":  ua.Name,
			"email":     ua.Email,
		},
	)

	if db.Error != nil {
		log.Fatal(db.Error)
		return u.Message(false, "An Error occurred")
	}
	ua.Password = ""
	response := u.Message(true, "User has been updated")
	response["account"] = ua
	return response
}