package handlers

import (
	"fmt"
	"net/http"
)

func identifyImage(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(0)
	fmt.Println(r.FormValue("delete"))
}
