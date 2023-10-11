package main

import (
	"crud/config"
	"crud/controllers/categorycontroller"
	"crud/controllers/homecontroller"
	"crud/controllers/productcontroller"
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
	http.HandleFunc("/products", productcontroller.Index)
	http.HandleFunc("/products/create", productcontroller.Create)
	http.HandleFunc("/products/detail", productcontroller.Create)
	http.HandleFunc("/products/update", productcontroller.Update)
	http.HandleFunc("/products/delete", productcontroller.Delete)

	log.Print("server running on port 8000")
	http.ListenAndServe(":8000", nil)
}
