package config

//Point of this package is to return DB which will help other files to interact with the DB

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// The connect func will helps us to open connection with our MYSQL DB
func Connect() {
	d, err := gorm.Open("mysql", "root:Shotgun@01@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local") // !st param is Type of DB 2nd Param is user:password/schemasName?
	//charset=utf8&parseTime=True&loc=Local is mandatory to be same
	if err != nil {
		panic(err)
	}
	db = d
}

//We will create a function which will use in other files
//WHich returns the db

func GetDB() *gorm.DB {
	return db
}
