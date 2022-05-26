package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Println("listening on port 3000...")
	log.Println("belajar gaes percobaan cicd 1")
	http.ListenAndServe(":3000", nil)
}
