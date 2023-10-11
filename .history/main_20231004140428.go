package main

import (
	"log"
	"net/http"
)

func main() {
	config Conne
	log.Print("server running on port 8000")
	http.ListenAndServe(":8000", nil)
}
