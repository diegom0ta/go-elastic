package api

import (
	"log"

	"github.com/olivere/elastic"
)

// Set the Elasticsearch URL
const url = "http://localhost:9200"

func ConnectToElasticsearch() *elastic.Client {

	// Create a new Elasticsearch client
	client, err := elastic.NewClient(elastic.SetURL(url))
	if err != nil {
		log.Fatal(err)
		return nil
	}

	// Return the client
	return client
}
