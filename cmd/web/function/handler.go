package function

import (
	"CS408-Golang-Hello-World-Full-Stack/cmd/web/pages"
	"CS408-Golang-Hello-World-Full-Stack/views"
	"log"
	"net/http"
	"regexp"
)

const ISE = "Internal Server Error"

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
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}

	if r.Header.Get("HX-Request") == "true" {
		err = views.PartialLoad(p).Render(r.Context(), w)
		if err != nil {
			log.Printf("Partial view failed to render")
			http.Error(w, ISE, http.StatusInternalServerError)
		}
		return
	}

	servedPage := views.View(p)
	err = servedPage.Render(r.Context(), w)

	if err != nil {
		log.Printf("page failed to render: %v", err)
		http.Error(w, ISE, http.StatusInternalServerError)
	}
}

// EditHandler : handles editing pages
func EditHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := pages.LoadPage(title)
	if err != nil {
		p = &pages.Page{Title: title, Body: []byte{}}
	}

	if r.Header.Get("HX-Request") == "true" {
		err = views.PartialEdit(p).Render(r.Context(), w)
		if err != nil {
			log.Printf("Partial Edit failed to render")
			http.Error(w, ISE, http.StatusInternalServerError)
		}
		return
	}
	servedPage := views.Edit(p)
	err = servedPage.Render(r.Context(), w)

	if err != nil {
		log.Printf("page failed to render: %v", err)
		http.Error(w, ISE, http.StatusInternalServerError)
	}
}

// SaveHandler : handles saving pages
func SaveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &pages.Page{Title: title, Body: []byte(body)}
	err := p.SavePage()
	if err != nil {
		log.Printf("failed to save page %q: %v", title, err)
		http.Error(w, ISE, http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

// RootHandler : handles root redirection
func RootHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, "/view/homepage", http.StatusFound)
}
