package main

import (
	"crud/config"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	// 1.Homepage
	http.HandleFunc("/")

	log.Print("server running on port 8000")
	http.ListenAndServe(":8000", nil)
}
