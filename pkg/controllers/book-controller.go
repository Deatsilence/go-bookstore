package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Deatsilence/go-bookstore/pkg/models"
	"github.com/Deatsilence/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, err := json.Marshal(newBooks)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("Error while parsing")
	}

	bookDetails, _ := models.GetBookById(ID)
	res, errMarsh := json.Marshal(bookDetails)

	if errMarsh != nil {
		http.Error(w, errMarsh.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook := &models.Book{}

	utils.ParseBody(r, createBook)
	b := createBook.CreateBook()

	res, err := json.Marshal(b)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("Error while parsing")
	}

	book := models.DeleteBook(ID)
	res, errMarsh := json.Marshal(book)

	if errMarsh != nil {
		http.Error(w, errMarsh.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("Error while parsing")
	}

	bookDetails, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)

	res, errMarsh := json.Marshal(bookDetails)

	if errMarsh != nil {
		http.Error(w, errMarsh.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
