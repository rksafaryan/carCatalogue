package handlers

import (
	"fmt"
	"net/http"
)

func Swag(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Swag")
}
