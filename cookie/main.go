package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "_cookie",
		Value:    "201709",
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
}
