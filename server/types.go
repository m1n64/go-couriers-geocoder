package server

import (
	"encoding/json"
	"net/http"
)

type Success struct {
	Success bool `json:"success"`
}

type Error struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func ErrorHandler(w http.ResponseWriter, err error) {
	var errorStruct = &Error{Success: false, Message: err.Error()}

	resp, _ := json.Marshal(errorStruct)

	w.WriteHeader(http.StatusInternalServerError)
	w.Write(resp)
}
