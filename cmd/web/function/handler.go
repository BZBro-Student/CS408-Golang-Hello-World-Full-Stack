package function

import (
	"CS408-Golang-Hello-World-Full-Stack/cmd/web/pages"
	"CS408-Golang-Hello-World-Full-Stack/views"
	"log"
	"net/http"
	"regexp"
)

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

// MakeHandler : wrapper for view edit and save handler functions
func MakeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

// ViewHandler : handles viewing of pages
func ViewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := pages.LoadPage(title)
	if err != nil {
		log.Printf("failed to LoadPage")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	servedPage := views.View(p)
	servedPage.Render(r.Context(), w)
}

// EditHandler : handles editing pages
func EditHandler(w http.ResponseWriter, r *http.Request, title string) {
	//TODO: implement
}

// SaveHandler : handles saving pages
func SaveHandler(w http.ResponseWriter, r *http.Request, title string) {
	//TODO: implement
}
