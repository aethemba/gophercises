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

	var tpl2 = `<html>
<body>
</body>
</html>`

	tables := []struct {
		t string
		c int
	}{
		{tpl, 1},
		{tpl2, 0},
	}

	for _, table := range tables {
		doc, _ := html.Parse(strings.NewReader(table.t))
		links := FindLinks(doc)
		if len(links) != table.c {
			t.Errorf("Expected %d, got %d", table.c, len(links))
		}
	}
}
