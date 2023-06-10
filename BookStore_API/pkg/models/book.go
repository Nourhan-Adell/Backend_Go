package models

// '''
// 	This is the Third step in which the controllers will call to connect to the database
// 	and do the SQL quires.
// '''

import (
	"github.com/Nourhan/go_crash_course/BookStore_API/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model //model: is something gives us a structure to help us store something in the database
	Name string `gorm: "json:"name""`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init (){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book{
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book{
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookByID(Id int64) (*Book, *gorm.DB){
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book{
	var book Book
	db.Where("ID=?", ID).Delete(&book)
	return book
}