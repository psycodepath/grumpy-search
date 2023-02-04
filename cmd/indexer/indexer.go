package main

import (
	"github.com/blevesearch/bleve/v2"
	"github.com/psycodepath/grumpy-search/pkg/index"
	grumpyMap "github.com/psycodepath/grumpy-search/pkg/sitemap"
	"log"
)

func main() {
	search, err := bleve.New("grumpy.bleve", index.CreateIndex())
	if err != nil {
		log.Fatal(err.Error())
	}
	sm, err := grumpyMap.GetSitemap("https://grumpy-learning.com/sitemap.xml")

	if err != nil {
		log.Fatal(err.Error())
	}

	for _, entry := range sm {
		if entry == nil {
			continue
		}
		p, err := index.GetPage(entry.GetLocation())
		if err != nil {
			log.Fatal(err.Error())
		}
		search.Index(p.Url, p)
	}

}
