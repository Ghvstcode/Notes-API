package utils

import (
	"encoding/json"
	"net/http"
)

//Message is exported
func Message(status bool, message string, code int) map[string]interface{} {
	return map[string]interface{} {"status" : status, "message" : message, "code" : code}
}

//Respond is exported
func RespondJson(w http.ResponseWriter, data map[string] interface{})  {
	w.Header().Add("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(data)
}