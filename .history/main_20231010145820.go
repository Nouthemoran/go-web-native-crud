package main

import (
	"crud/config"
	"crud/controllers/categorycontroller"
	"crud/controllers/homecontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	// 1.Homepage
	http.HandleFunc("/", homecontroller.Welcome)

	// 2.Categories
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/create", categorycontroller.Create)
	http.HandleFunc("/categories/update", categorycontroller.Update)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	// 3.Products
	http.Hand

	log.Print("server running on port 8000")
	http.ListenAndServe(":8000", nil)
}
