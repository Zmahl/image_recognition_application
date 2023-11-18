package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	receiveImageHandler := http.HandlerFunc(receiveImage)
	http.Handle("/identify-image", receiveImageHandler)
	http.ListenAndServe(":8080", nil)
}

func receiveImage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	file, h, err := r.FormFile("photo")

	tmpfile, err := os.Create("./images" + h.Filename)
	defer tmpfile.Close()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(tmpfile, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	return
}
