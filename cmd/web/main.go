package main

import (
	"CS408-Golang-Hello-World-Full-Stack/cmd/web/function"
	"log"
	"net/http"
)

func startServer(port string) {
	log.Printf("Server starting on http://localhost%s\n", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", function.RootHandler)
	http.HandleFunc("/view/", function.MakeHandler(function.ViewHandler))
	http.HandleFunc("/edit/", function.MakeHandler(function.EditHandler))
	http.HandleFunc("/save/", function.MakeHandler(function.SaveHandler))

	port := ":8080"
	startServer(port)

}
