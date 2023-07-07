package routes

import (
	"github.com/gorilla/mux"
	"github.com/guilledipa/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(r *mux.Router) {
	r.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	r.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/book/{bookID}", controllers.GetBookByID).Methods("GET")
	r.HandleFunc("/book/{bookID}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/book/{bookID}", controllers.DeleteBook).Methods("DELETE")
}
