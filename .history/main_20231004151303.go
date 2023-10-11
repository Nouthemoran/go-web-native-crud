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
	http.HandleFunc("/", homecontroller.Welcome)

	log.Print("server running on port 8000")
	http.ListenAndServe(":8000", nil)
}
