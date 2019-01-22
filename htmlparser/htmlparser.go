package htmlparser

import (
	"io"

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

func BuildLink(n *html.Node) Link {
	for _, a := range n.Attr {
		if a.Key == "href" {
			return Link{a.Val, n.FirstChild.Data}
		}
	}
	return Link{}
}

func FindLinks(n *html.Node) []Link {

	var res []Link

	isLink := ParseLink(n)

	if isLink == true {
		l := BuildLink(n)
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
