package controllers

import (
	"encoding/json"
	"gorilla/utils"
	"net/http"
)

func ListBookController(w http.ResponseWriter, req *http.Request) {
	books := utils.FindBooks()
	jsonResponse, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if len(books) == 0 {
		result := []string{}
		jsonResponse, _ := json.Marshal(result)
		w.Write(jsonResponse)
	} else {
		w.Write(jsonResponse)
	}
}
