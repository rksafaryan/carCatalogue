package handlers

import (
	"fmt"
	"github.com/pborman/uuid"
	"net/http"
)

func SwaggerYaml(w http.ResponseWriter, r *http.Request) {
	uuidWithHyphen := uuid.NewRandom()
	fmt.Fprintln(w, "swagger yaml"+string(uuidWithHyphen))
}
