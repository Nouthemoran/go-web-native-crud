package main

import (
	"log"
	"net/http"
)

func main() {
	log.Print("")
	http.ListenAndServe(":8000", nil)
}