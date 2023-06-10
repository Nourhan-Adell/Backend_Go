package controllers

// '''
// 	In this file we will do the functions which is requested to do by the user

// 	This is the second step --> then goes to models
// '''

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Nourhan/go_crash_course/BookStore_API/pkg/models"
	"github.com/Nourhan/go_crash_course/BookStore_API/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book


func GetBooks(w http.ResponseWriter, r *http.Request){
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)	// Convert it to json
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK) // the Status will be 200
	w.Write(res)	// Send the response to the frontend
}

func GetBookByID(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars ["bookId"] 
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookByID(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK) // the Status will be 200
	w.Write(res)	// Send the response to the frontend
}

func CreateBook(w http.ResponseWriter, r *http.Request){
	createBook := &models.Book{}
	utils.PasreBody(r, createBook)
	b := createBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK) // the Status will be 200
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookID := vars ["bookID"] 
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)  
	w.Write(res)	
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	var updateBook = &models.Book{}
	utils.PasreBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	bookDetails, db := models.GetBookByID(ID)

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
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK) 
	w.Write(res)
}