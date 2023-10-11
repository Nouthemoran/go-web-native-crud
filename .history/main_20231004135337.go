package main

import (
	"log"
	"net/http"
)

func main() {
	log.Print("server running on port ")
	http.ListenAndServe(":8000", nil)
}