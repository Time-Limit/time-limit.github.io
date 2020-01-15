package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./blog/dist")))
	http.ListenAndServe(":5555", nil)
}
