package main

import (
	"log"
	"net/http"
)

func main() {
	config C
	log.Print("server running on port 8000")
	http.ListenAndServe(":8000", nil)
}
