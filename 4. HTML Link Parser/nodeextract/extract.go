package nodeextract

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

type htmlLink struct {
	link string
	text string
}

// <a href="https://www.example.com">Visit Example</a>
// 'a' is tag/data and 'href' is attribute

// Extract a link node (<a href /> tags) and its text node
// This will be called for different nodes
func getLink(n *html.Node) htmlLink {
	var link htmlLink
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			link.link = attr.Val
			link.text = extractText(n)
			break // I don't want to loop on more attributes
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		getLink(c)
	}

	return link
}

// Extract the text nodes
func extractText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}

	var sb strings.Builder
	// DFS
	for c := n.FirstChild; c != nil; c = n.NextSibling {
		sb.WriteString(extractText(c) + " ")
	}
	fmt.Println("Texts: ", sb)
	return sb.String()
}

// Extract the pointer to href nodes
func ExtractHref(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}

	// Recursive DFS Call
	var hrefs []*html.Node
	for c := n.FirstChild; c != nil; c = n.NextSibling {
		hrefs = append(hrefs, ExtractHref(c)...)
	}

	return hrefs
}
