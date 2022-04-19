package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gorilla/utils"
	"net/http"
)

func ReadBookController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	book, isBookExists := utils.FindByTitleBook(title)
	if isBookExists {
		response, _ := json.Marshal(book)

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	} else {
		type notFound struct {
			Msg string `json:"msg"`
		}

		result := notFound{Msg: "NOTFOUND"}

		w.WriteHeader(http.StatusNotFound)
		response, _ := json.Marshal(result)
		w.Write(response)
	}
}
