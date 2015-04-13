package esql

import (
	"testing"

	. "github.com/otiai10/mint"
)

func TestHoge(t *testing.T) {
	q := Search("index_foo/type_bar").Query(
		Filtered(
			Filter(
				And(
					Range(),
					Term(map[string]interface{}{
						"foo":  "bar",
						"hoge": "fuga",
					}),
				),
			),
			Query(),
		),
	)
	Expect(t, q.String()).ToBe(`{"query":{"filtered":{"filter":{"and":{"range":{},"term":{"foo":"bar","hoge":"fuga"}}},"query":{}}}}`)
}
