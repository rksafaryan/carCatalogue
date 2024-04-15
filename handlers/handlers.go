package handlers

import (
	"fmt"
	"net/http"
)

func SwaggerYaml(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "swagger yaml")
}
