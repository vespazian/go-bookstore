package routes

import (
	"github.com/gorilla/mux"
	"github.com/vespazian/go-bookstore/controllers"
)

var RegisterBookStore = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook)
}
