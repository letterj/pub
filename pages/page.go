package pages

import (
	"io/ioutil"

	"github.com/russross/blackfriday"
)

// Page is the Web page structure
type Page struct {
	Title  string
	Author string
	HTML   string
	File   string
}

// NewPage is the constructor for Page
func NewPage(filename string) (*Page, error) {
	p := &Page{File: filename}
	return p, p.load()
}

// load Method to load details about a page
func (p *Page) load() error {
	c, err := LoadConfig("config.json")
	if err != nil {
		return err
	}

	data, err := ioutil.ReadFile(p.File)
	if err != nil {
		return err
	}

	p.HTML = string(blackfriday.MarkdownCommon(data))
	p.Title = c.Name()
	p.Author = c.Author()

	return nil
}
