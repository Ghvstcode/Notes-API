package models

import (
	"github.com/jinzhu/gorm"

	u "github.com/GhvstCode/Notes-API/utils"
)

//a struct to represent Notes
type Note struct {
	gorm.Model
	Title string `json:"title"`
	Content   string `json:"content"`
	UserID uint `json:"user_id"`
}

func (n *Note) Validate()(map[string]interface{}, bool){
	if len(n.Title) < 0 {
		return u.Message(false, "Note must contain title"), false
	}

	if len(n.Content) < 0 {
		return u.Message(false, "Note must contain content"), false
	}

	if n.UserID <= 0 {
		return u.Message(false, "unauthorized access"), false
	}
	return u.Message(true, "Validated"), true
}

func (n *Note) Create() map[string]interface{} {

	if resp, ok := n.Validate(); !ok {
		return resp
	}

	GetDB().Create(n)

	resp := u.Message(true, "success")
	resp["note"] = n
	return resp

}
