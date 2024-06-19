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
	_ = doc
	// get the href nodes
	ExtractHref(doc)

}
