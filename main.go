package main

import (
	"github.com/rksafaryan/carcatalogue/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.SwaggerYaml)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
