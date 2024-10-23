package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"asciiWeb/handler"
)

func init() {
	var err error
	handler.Templates, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fileServer := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fileServer))
	http.HandleFunc("/", handler.HandleMainPage)
	http.HandleFunc("/ascii-art", handler.HandleAsciiArt)
	fmt.Println("Server starting at http://localhost:6500")
	log.Fatal(http.ListenAndServe(":6500", nil))
}
