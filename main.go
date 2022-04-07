package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type booksType struct {
	Title string `json:"title"`
	Stock int    `json:"stock"`
}

var bookStore = []booksType{}
var book = make(map[string]booksType)

func CheckBookExists(title string) bool {
	result := false
	for _, b := range bookStore {
		if b.Title == title {
			result = true
		}
	}
	return result
}

func ReadBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]

	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
}

func CreateBook(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	title := vars["title"]

	var input booksType
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Decode error! please check your JSON formating.")
		return
	}

	isBookExist := CheckBookExists(title)
	if !isBookExist {
		bookStore = append(bookStore, booksType{Title: title, Stock: input.Stock})
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "Created Book")
	} else {
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintf(w, "Book is Already Exists")
	}
}

func DeleteBook(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	title := vars["title"]

	arr := []booksType{}
	isDelete := false
	for index, b := range bookStore {
		if b.Title != title {
			arr = append(arr, bookStore[index])
		} else {
			isDelete = true
		}
	}

	bookStore = arr

	if isDelete {
		fmt.Fprintf(w, "Delete Book")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Not Found Book.")
	}
}

func UpdateStock(title string, stock int) bool {
	result := false
	for i, b := range bookStore {
		if b.Title == title {
			bookStore[i].Stock = stock
			result = true
		}
	}

	return result
}

func UpdateBook(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	title := vars["title"]

	var input booksType
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Decode error! please check your JSON formating.")
		return
	}

	isUpdateStock := UpdateStock(title, input.Stock)

	if isUpdateStock {
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, "Update Book")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Not Found Book.")
	}

}

func ListBook(w http.ResponseWriter, req *http.Request) {
	foo_marshalled, _ := json.Marshal(bookStore)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(foo_marshalled)) // write response to ResponseWriter (w)
}

func main() {
	port := 8080
	listenPort := fmt.Sprintf(":%d", port)
	log.Printf("App listen on port %d\n", port)

	router := mux.NewRouter()

	router.HandleFunc("/books", ListBook).Methods("GET")
	router.HandleFunc("/books/{title}/page/{page}", ReadBook).Methods("GET")
	router.HandleFunc("/books/{title}", CreateBook).Methods("POST")
	router.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")

	http.ListenAndServe(listenPort, router)
}
