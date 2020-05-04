package utils

import (
	"encoding/json"
	"net/http"
)

func Message(status int, message string) (map[string] interface{}) {
	return map[string]interface{} {"status" : status, "message" : message}
}

func Respond(w http.ResponseWriter, data map[string] interface{})  {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(data["status"].(int))
	delete(data, "status")
	json.NewEncoder(w).Encode(data)
}
