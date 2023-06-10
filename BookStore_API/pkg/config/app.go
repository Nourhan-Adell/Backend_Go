package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db * gorm.DB
)

// d: is the variable to connect the database
// Connect with the database
func Connect(){
	d, err := gorm.Open("mysql", "akhil:Axlesharma@tcp(sql.hosting.com:3306)/simplerest?charset=utf8&parseTime=True&loc=Local")	//username:password/tables name
	if err != nil{
		panic(err)
	}
	db = d	
}

func GetDB() *gorm.DB{
	return db
}