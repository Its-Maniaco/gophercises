package main

import (
	"flag"
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	fileName := flag.String("html", "ex1.html", "Name of the html to parse")
	flag.Parse()

	// Open html file
	file, err := os.Open(*fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Parse html
	doc, err := html.Parse(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	// <a href="https://www.example.com">Visit Example</a>
	// 'a' is tag/data and 'href' is attribute
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					fmt.Print(attr.Val + ": ")
					for c := n.FirstChild; c != nil; c = c.NextSibling {
						if c.Type == html.TextNode {
							fmt.Println(c.Data + " ")
						}
					}
					break // I don't want to loop on more attributes
				}
			}
		}
		// Recursive Call
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}
