package routes

import (
	"github.com/gorilla/mux"
	"github.com/ngollach/go-bookstore/pkg/controllers"
)

//Creating a variable RegisterBookStore to have all my routes
//Which helps to get control to my controller where I meet my logic

var RegisterBooksStore = func(router *mux.Router) {

	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
