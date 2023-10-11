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
	http.HandleFunc("/categories", homecontroller.Index)
	http.HandleFunc("/categories/add", homecontroller.Create)
	http.HandleFunc("/categories", homecontroller.Update)
	http.HandleFunc("/categories", homecontroller.)

	log.Print("server running on port 8000")
	http.ListenAndServe(":8000", nil)
}
