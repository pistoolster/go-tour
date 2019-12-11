package main

import (
	"net/http"
)

func ServeFileFunc(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./test_xml.png")
}

func main() {
	http.HandleFunc("/test", ServeFileFunc)
	http.ListenAndServe(":8080", nil)
}
