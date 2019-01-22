package htmlparser

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func ParseLink(n *html.Node) bool {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				return true
			}
		}
	}
	return false
}

func text(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}

	if n.Type != html.ElementNode {
		return ""
	}

	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += text(c) + " "
	}
	return strings.Join(strings.Fields(ret), " ")
}

func BuildLink(n *html.Node) (Link, error) {
	for _, a := range n.Attr {
		if a.Key == "href" {
			return Link{a.Val, text(n)}, nil
		}
	}
	return Link{}, fmt.Errorf("Node is not a link")
}

func FindLinks(n *html.Node) []Link {

	var res []Link

	isLink := ParseLink(n)

	if isLink == true {
		l, err := BuildLink(n)
		if err != nil {
			panic(err)
		}
		res = append(res, l)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		res = append(res, FindLinks(c)...)

	}
	return res
}

func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)

	if err != nil {
		return nil, err
	}

	links := FindLinks(doc)
	return links, nil
}
