package controllers

import (
	"encoding/json"
	"go-book-store/pkg/config"
	"go-book-store/pkg/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	config.DB.Find(&books)
	json.NewEncoder(w).Encode(books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var book models.Book
	config.DB.First(&book, id)
	json.NewEncoder(w).Encode(book)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)
	config.DB.Create(&book)
	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var book models.Book
	config.DB.First(&book, id)
	json.NewDecoder(r.Body).Decode(&book)
	config.DB.Save(&book)
	json.NewEncoder(w).Encode(book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var book models.Book
	config.DB.Delete(&book, id)
	json.NewEncoder(w).Encode("Book deleted successfully")
}
