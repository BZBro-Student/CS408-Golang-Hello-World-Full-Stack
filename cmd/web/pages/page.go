package pages

import (
	"log"
	"os"
)

const dataPath string = "cmd/web/data/"

type Page struct {
	Title string
	Body  []byte
}

// savePage : saves contents of page, returns error if write fails and nil otherwise
func (p *Page) SavePage() error {
	err := os.WriteFile(dataPath+p.Title+".txt", p.Body, 0644)
	if err != nil {
		log.Printf("Failed page save")
		return err
	}
	return nil
}

// loadPage : loads page contents into a page struct, return error on fail and nil otherwise
func LoadPage(title string) (*Page, error) {
	content, err := os.ReadFile(dataPath + title + ".txt")

	if err != nil {
		log.Printf("Failed page load")
		return nil, err
	}
	page := Page{Title: title, Body: content}
	return &page, err
}
