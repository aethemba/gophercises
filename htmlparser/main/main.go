package main

import (
	"fmt"
	"gophercises/htmlparser"
	"strings"

	"golang.org/x/net/html"
)

var tpl = `<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>
	</head>
	<body>
		<h1>{{.Title}}</h1>
		<a href="/linkie">For the fun of links!</a>
        <div class="story">
            {{range .Story}}
            <p>{{.}}</p>
            {{end}}
        </div>
        <h3>Options</h3>
        <ul class="options">
            {{range .Options}}
            <li>{{.Text}}<span> <a href="/stories/{{.Arc}}">Select</a> </span></li>
            {{else}}
            <li><a href="/stories/intro">Start again</a></li>
            {{end}}
        </ul>
	</body>
</html>
`

func main() {
	doc, err := html.Parse(strings.NewReader(tpl))

	if err != nil {
		fmt.Printf("%s", err)
	}

	links := htmlparser.FindLinks(doc)

	fmt.Printf("\nFound the following links: %#v\n", links)

}
