package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"gorilla/utils"
	"net/http"
)

func DeleteBookController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	isDelete := utils.DeleteBook(title)

	if isDelete {
		fmt.Fprintf(w, "Delete Book")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Not Found Book.")
	}
}
