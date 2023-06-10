package main

import (
	"log"
	"net/http"

	"github.com/Nourhan/go_crash_course/BookStore_API/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main(){
	r := mux.NewRouter()
	routes.RegisterBookStore(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:7000",r))
}