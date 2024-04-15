package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/rksafaryan/carcatalogue/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", SwaggerYaml)
	http.HandleFunc("/swag", handlers.Swag)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func SwaggerYaml(w http.ResponseWriter, r *http.Request) {
	uuidWithHyphen, _ := uuid.NewRandom()
	fmt.Fprintln(w, "swagger yaml "+uuidWithHyphen.String())
}
