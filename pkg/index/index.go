package index

import (
	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/gocolly/colly"
)

// Page represents a page on a website
type Page struct {
	Url     string `json:"url"`
	Title   string `json:"title"`
	Robots  string `json:"robots"`
	Content string `json:"content"`
}

// Type returns the type used by bleve to determine which document mapping to apply
func (p *Page) Type() string {
	return "Page"
}

// GetPage grabs a page by url and maps it to the Page struct
//
// colly is used so not much of our spare braincells are needed
func GetPage(url string) (Page, error) {
	var title, content, robots string

	title = ""
	content = ""
	robots = ""
	c := colly.NewCollector()

	c.OnHTML("title", func(e *colly.HTMLElement) {
		title = e.Text
	})

	c.OnHTML("meta[name=\"robots\"]", func(e *colly.HTMLElement) {
		robots = e.Attr("content")
	})

	c.OnHTML("div.e-content", func(e *colly.HTMLElement) {
		content = e.Text
	})

	c.Visit(url)

	return Page{Url: url, Title: title, Robots: robots, Content: content}, nil
}

// CreatePageMapping creates the document mapping used by bleve for pages
func CreatePageMapping() *mapping.DocumentMapping {
	pageMapping := bleve.NewDocumentMapping()

	urlMapping := bleve.NewTextFieldMapping()
	urlMapping.Analyzer = "en"

	titleMapping := bleve.NewTextFieldMapping()
	titleMapping.Analyzer = "en"

	contentMapping := bleve.NewTextFieldMapping()
	contentMapping.Analyzer = "en"

	pageMapping.AddFieldMappingsAt("url", urlMapping)
	pageMapping.AddFieldMappingsAt("title", titleMapping)
	pageMapping.AddFieldMappingsAt("content", titleMapping)

	return pageMapping
}

func CreateIndex() *mapping.IndexMappingImpl {
	index := bleve.NewIndexMapping()
	index.AddDocumentMapping("Page", CreatePageMapping())
	return index
}
