package routes

// '''
// 	In this file we start with udentify the needed routes in our API

// 	This is the first step --> Then goes to the Controllers
// '''
import (
	"github.com/Nourhan/go_crash_course/BookStore_API/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStore = func (router *mux.Router)  {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookByID).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}