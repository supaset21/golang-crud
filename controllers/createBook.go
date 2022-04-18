package controllers

import (
	"encoding/json"
	"fmt"
	"gorilla/utils"
	"net/http"

	"gorilla/models"

	"github.com/gorilla/mux"
)

func CreateBookController(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	title := vars["title"]

	var input models.BooksType
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Decode error! please check your JSON formating.")
		return
	}

	isBookExist := utils.CheckBookExists(title)
	if !isBookExist {
		utils.CreateBook(title, input.Stock)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "Created Book")
	} else {
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintf(w, "Book is Already Exists")
	}
}
