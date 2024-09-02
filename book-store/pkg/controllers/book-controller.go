package controllers

import (
	"book-store/pkg/config"
	"book-store/pkg/errors"
	"book-store/pkg/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getBoodById(r *http.Request) (*models.Book, error) {
	book := new(models.Book)

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		return nil, &errors.BadRequest{Msg: "invalid book id"}
	}

	result := config.DB.First(book, id)
	if result.Error != nil {
		return nil, &errors.NotFound{Msg: "book not found"}
	}

	return book, nil
}

func GetBooks(r *http.Request) (*models.Response, error) {
	var books []models.Book
	config.DB.Find(&books)

	return &models.Response{
		StatusCode: http.StatusOK,
		Body:       books,
	}, nil
}

func GetBook(r *http.Request) (*models.Response, error) {
	book, err := getBoodById(r)
	if err != nil {
		return nil, err
	}

	return &models.Response{
		StatusCode: http.StatusOK,
		Body:       book,
	}, nil
}

func CreateBook(r *http.Request) (*models.Response, error) {
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)
	config.DB.Create(&book)

	return &models.Response{
		StatusCode: http.StatusCreated,
		Body:       book,
	}, nil
}

func UpdateBook(r *http.Request) (*models.Response, error) {
	book, err := getBoodById(r)
	if err != nil {
		return nil, err
	}

	var updatedBook models.Book
	json.NewDecoder(r.Body).Decode(&updatedBook)

	book.Author = updatedBook.Author
	book.ISBN = updatedBook.ISBN
	book.Title = updatedBook.Title
	config.DB.Save(book)

	return &models.Response{
		StatusCode: http.StatusOK,
		Body:       book,
	}, nil
}

func DeleteBook(r *http.Request) (*models.Response, error) {
	book, err := getBoodById(r)
	if err != nil {
		return nil, err
	}
	config.DB.Delete(book)

	return &models.Response{
		StatusCode: http.StatusNoContent,
		Body:       "Book deleted successfully",
	}, nil
}
