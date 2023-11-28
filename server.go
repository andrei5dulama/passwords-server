package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8800", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		file, handler, err := r.FormFile("passwordList")
		if err != nil {
			fmt.Println("Error uploading file:", err)
			return
		}
		defer file.Close()

		f, err := os.Create(handler.Filename)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer f.Close()

		io.Copy(f, file)

		fmt.Fprintf(w, "File uploaded successfully!")
		return
	}

	fmt.Fprintf(w, "Passwords server")
}
