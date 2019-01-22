package htmlparser

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestFindLinks(t *testing.T) {
	var tpl = `<html>
<body>
  <a href="/dog-cat">dog cat <!-- commented text SHOULD NOT be included! --></a>
</body>
</html>`

	doc, _ := html.Parse(strings.NewReader(tpl))
	links := FindLinks(doc)
	if len(links) != 1 {
		t.Error("Expected 1, got ", len(links))
	}
}
