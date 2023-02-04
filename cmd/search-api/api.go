package main

import (
	"encoding/json"
	"github.com/blevesearch/bleve/v2"
	"log"
	"net/http"
)

func main() {

	index, err := bleve.Open("grumpy.bleve")

	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		params := request.URL.Query()
		query := bleve.NewMatchQuery(params.Get("search"))
		result, err := index.Search(bleve.NewSearchRequestOptions(query, 25, 0, false))
		if err != nil {
			writer.WriteHeader(500)
		}
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(200)

		searchResults, _ := json.Marshal(result)
		_, err = writer.Write(searchResults)

		if err != nil {
			log.Fatal(err)
		}
	})

	http.ListenAndServe(":1337", nil)
}
