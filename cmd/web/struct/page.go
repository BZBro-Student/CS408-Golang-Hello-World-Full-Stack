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
func (p *Page) savePage() error {
	err := os.WriteFile("cmd/web/data/"+p.title+".txt", p.body, 0644)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// loadPage : loads page contents into a page struct, return error on fail and nil otherwise
func loadPage(title string) (*Page, error) {

}
