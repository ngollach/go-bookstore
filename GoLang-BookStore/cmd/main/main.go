package main

//Here we will create a server
//We will also define our localHost
//We will tell our GoLang where our router resides i.e, Bookstore-rotes has the files

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ngollach/go-bookstore/pkg/routes"
)

func main() {

	r := mux.NewRouter()

	routes.RegisterBooksStore(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", r)) //ListenAndServe is fucntion which helps us to create a server we pass our port or address
	//which is inside the http
}
