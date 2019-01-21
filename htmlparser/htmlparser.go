package htmlparser

import (
	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func FindLinks(n *html.Node) []Link {
	var l Link
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				l.Href = a.Val
				break
			}
		}
		l.Text = n.FirstChild.Data
	}

	var res []Link
	if l.Href != "" {
		res = append(res, l)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result := FindLinks(c)
		res = append(res, result...)

	}
	return res
}
