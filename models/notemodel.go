package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"

	u "github.com/GhvstCode/Notes-API/utils"
)

//a struct to represent Notes
type Note struct {
	gorm.Model
	ID        uint `gorm:"primary_key"`
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
		resp = u.Message(false, "An Error occurred!")
		return resp
	}

	GetDB().Create(n)

	resp := u.Message(true, "success")
	resp["note"] = n
	return resp

}
//
func GetNotes(id uint) []*Note {
	note := make([]*Note, 0)
	err := GetDB().Table("notes").Where("user_id = ?", id).Find(&note).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return note
}

func (n *Note)Update(noteID uint) map[string]interface{} {
	if resp, ok := n.Validate(); !ok {
		resp = u.Message(false, "An Error occurred!")
		return resp
	}
	db = GetDB().Debug().Model(&Note{}).Where("ID = ?", noteID).Take(&Note{}).UpdateColumns(
		map[string]interface{}{
			"ID": noteID,
			"title":  n.Title,
			"content":  n.Content,
		},
	)
	n.ID = noteID
	if db.Error != nil {
		log.Fatal(db.Error)
		return u.Message(false, "An Error occurred")
	}
	response := u.Message(true, "User has been updated")
	response["account"] = n
	return response
}

func (n *Note) DeleteNote(uid uint32) map[string]interface{} {
	db = GetDB().Debug().Model(&Note{}).Where("id = ?", uid).Take(&Note{}).Delete(&Note{})

	if db.Error != nil {
		return u.Message(true, "Note has been deleted")
	}

	response := u.Message(true, "Note has been deleted")
	response["Data"] = &Note{}
	return response
}