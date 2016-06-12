package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fs := http.FileServer(http.Dir("static"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		reader, err := os.Open("index.html")
		if err != nil {
			w.Write([]byte("error reading index.html!!!"))
		}

		io.Copy(w, reader)
	})

	fmt.Println("Listening on port :8080")
	http.ListenAndServe(":8080", nil)
}
