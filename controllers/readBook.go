package controllers

import (
	"fmt"
	"net/http"
)

func ReadBook(w http.ResponseWriter, r *http.Request) bool {
	fmt.Println("HELLO")
	return true
}
