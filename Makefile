.phony: index run

index:
	go run github.com/psycodepath/grumpy-search/cmd/indexer

build:
	go build github.com/psycodepath/grumpy-search/cmd/search-api

run: build
	./search-api