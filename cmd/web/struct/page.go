package web

import (
	"log"
	"os"
)

const dataPath string = "cmd/web/data/"

type Page struct {
	title string
	body  []byte
}

// savePage : saves contents of page, returns error if write fails and nil otherwise
func (p *Page) SavePage() error {
	err := os.WriteFile(dataPath+p.title+".txt", p.body, 0644)
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
	page := Page{title: title, body: content}
	return &page, err

}
