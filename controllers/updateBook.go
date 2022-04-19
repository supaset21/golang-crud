package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gorilla/models"
	"gorilla/utils"
	"net/http"
)

func UpdateBookController(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	title := vars["title"]

	var input models.BooksType
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Decode error! please check your JSON formating.")
		return
	}

	isUpdateStock := utils.UpdateStock(title, input.Stock)

	if isUpdateStock {
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, "Update Book")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Not Found Book.")
	}
}
