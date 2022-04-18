package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	C "gorilla/controllers"

	"github.com/gorilla/mux"
)

type booksType struct {
	Title string `json:"title"`
	Stock int    `json:"stock"`
}

var bookStore = []booksType{}
var book = make(map[string]booksType)

func ReadBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	result := booksType{}
	isBook := false

	for index, b := range bookStore {
		if b.Title == title {
			result = bookStore[index]
			isBook = true
		}
	}

	if isBook {
		response, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(response))
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Not Found Book.")
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

//func ListBook(w http.ResponseWriter, req *http.Request) {
//	foo_marshalled, _ := json.Marshal(bookStore)
//	w.WriteHeader(http.StatusOK)
//	fmt.Fprint(w, string(foo_marshalled)) // write response to ResponseWriter (w)
//}

func main() {
	port := 8080
	listenPort := fmt.Sprintf(":%d", port)
	log.Printf("App listen on port %d\n", port)

	router := mux.NewRouter()

	router.HandleFunc("/books", C.ListBookController).Methods("GET")
	router.HandleFunc("/books/{title}", ReadBook).Methods("GET")
	router.HandleFunc("/books/{title}", C.CreateBookController).Methods("POST")
	router.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")

	http.ListenAndServe(listenPort, router)
}
