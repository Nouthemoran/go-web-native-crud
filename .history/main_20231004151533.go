package main

import (
	"crud/config"
	"crud/controllers/homecontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	// 1.Homepage
	http.HandleFunc("/", homecontroller.Welcome)

	// 2.Categories
	http.HandleFunc("/categories", category.Index)
	http.HandleFunc("/categories/create", category.Create)
	http.HandleFunc("/categories/update", category.Update)
	http.HandleFunc("/categories/delete", category.Delete)

	log.Print("server running on port 8000")
	http.ListenAndServe(":8000", nil)
}
