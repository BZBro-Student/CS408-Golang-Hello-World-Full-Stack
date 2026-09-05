package main

import (
	"CS408-Golang-Hello-World-Full-Stack/cmd/web/function"
	"net/http"
)

func main() {
	http.HandleFunc("/view/", function.MakeHandler(function.ViewHandler))
}
