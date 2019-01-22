package htmlparser

import (
	"strings"
	"testing"
)

func TestParse(t *testing.T) {

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
		r := strings.NewReader(table.t)

		links, _ := Parse(r)

		if len(links) != table.c {
			t.Errorf("Expected %d, got %d", table.c, len(links))
		}
	}
}
