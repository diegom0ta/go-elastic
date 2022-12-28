package handler

import (
	"context"
	"fmt"

	"github.com/olivere/elastic"
	"golang.org/x/net/html"
)

type Product struct {
	Name  string
	Price string
}

func indexProduct(product *Product, client *elastic.Client) {
	// Index the product in Elasticsearch
	_, err := client.Index().
		Index("products").
		Type("product").
		Id(product.Name).
		BodyJson(product).
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print a message to the console
	fmt.Println("Product indexed:", product.Name)
}

func ExtractProductData(n *html.Node, client *elastic.Client) {
	// Check if the node is an element node
	if n.Type == html.ElementNode {
		// Check if the element is a "div" with the class "product"
		if n.Data == "div" {
			for _, attr := range n.Attr {
				if attr.Key == "class" && attr.Val == "product" {
					// Extract the product data from the element
					product := extractProductDataFromElement(n)

					// Index the product data in Elasticsearch
					indexProduct(product, client)
				}
			}
		}
	}

	// Recursively call the function for each child node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ExtractProductData(c, client)
	}
}

func extractProductDataFromElement(n *html.Node) *Product {
	// Initialize a product struct
	product := &Product{}

	// Traverse the children of the element
	for c := n.FirstChild; c != nil; {
		// Check if the child is an element node
		if c.Type == html.ElementNode {
			// Extract the product name
			if c.Data == "h3" {
				for _, attr := range c.Attr {
					if attr.Key == "class" && attr.Val == "name" {
						product.Name = c.FirstChild.Data
					}
				}
			}
			// Extract the product price
			if c.Data == "span" {
				for _, attr := range c.Attr {
					if attr.Key == "class" && attr.Val == "price" {
						product.Price = c.FirstChild.Data
					}
				}
			}
		}
	}
	// Return the product struct
	return product
}
