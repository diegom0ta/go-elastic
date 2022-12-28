package main

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/diegom0ta/go-elastic/api"
	"github.com/diegom0ta/go-elastic/handler"
	"golang.org/x/net/html"
)

// Set the starting URL
const url = "https://www.amazon.com/products"

func main() {

	// Make an HTTP GET request to the URL
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Parse the HTML content
	doc, err := html.Parse(strings.NewReader(string(body)))
	if err != nil {
		log.Fatal(err)
		return
	}

	// Connect to Elasticsearch
	client := api.ConnectToElasticsearch()
	if client == nil {
		log.Print("Unable to connect to Elasticsearch")
		return
	}

	// Call the function to extract product data from the HTML
	handler.ExtractProductData(doc, client)
}
