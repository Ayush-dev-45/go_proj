package main

import (
	"bookstore-management/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.BookstoreRoutes(r)
	http.Handle("/", r);
	log.Fatal(http.ListenAndServe(":9090", r))
}