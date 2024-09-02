package controllers

import (
	"book-store/pkg/config"
	"book-store/pkg/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	config.DB.Find(&books)
	json.NewEncoder(w).Encode(books)
}

func GetBook(w http.ResponseWriter, r *http.Request) *models.Response {
	response := new(models.Response)
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		response = &models.Response{
			StatusCode: http.StatusBadRequest,
			Body:       "Invalid book ID",
		}
		return response
	}

	var book models.Book
	result := config.DB.First(&book, id)
	if result.Error != nil {
		response = &models.Response{
			StatusCode: http.StatusNotFound,
			Body:       "Book not found",
		}
	} else {
		response = &models.Response{
			StatusCode: http.StatusOK,
			Body:       book,
		}
	}

	return response
}

func CreateBook(w http.ResponseWriter, r *http.Request) *models.Response {
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)
	config.DB.Create(&book)
	return &models.Response{
		StatusCode: http.StatusCreated,
		Body:       book,
	}
}

func UpdateBook(w http.ResponseWriter, r *http.Request) *models.Response {
	response := new(models.Response)
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		response = &models.Response{
			StatusCode: http.StatusBadRequest,
			Body:       "Invalid book ID",
		}
		return response
	}

	var book models.Book
	config.DB.First(&book, id)
	json.NewDecoder(r.Body).Decode(&book)
	config.DB.Save(&book)

	response = &models.Response{
		StatusCode: http.StatusOK,
		Body:       book,
	}
	return response
}

func DeleteBook(w http.ResponseWriter, r *http.Request) *models.Response {
	response := new(models.Response)
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		response = &models.Response{
			StatusCode: http.StatusBadRequest,
			Body:       "Invalid book ID",
		}
		return response
	}

	var book models.Book
	config.DB.Delete(&book, id)

	response = &models.Response{
		StatusCode: http.StatusNoContent,
		Body:       "Book deleted successfully",
	}
	return response
}
