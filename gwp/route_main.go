package main

import (
	"net/http"
)

func index(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, "你好，世界！", "index")
}
