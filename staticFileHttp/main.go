package main

import "net/http"

func main() {
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":8080", nil)
}
