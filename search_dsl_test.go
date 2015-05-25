package escalgo

import (
	"encoding/json"
	"testing"

	. "github.com/otiai10/mint"
)

func TestSearchDSL(t *testing.T) {
	search := Search("twitter/tweet")
	Expect(t, search).TypeOf("*escalgo.SearchDSL")
	b, err := json.Marshal(search)
	Expect(t, err).ToBe(nil)
	Expect(t, string(b)).ToBe(`{}`)

	search.Set([]Sortable{Sort{"hoge", true}, Sort{"fuga", false}})
	b, err = json.Marshal(search)
	Expect(t, err).ToBe(nil)
	Expect(t, string(b)).ToBe(`{"sort":[{"hoge":"desc"},"fuga"]}`)

	filtered := &Filtered{}
	filtered.Query().Bool().Should(Terms{"lang", []string{"go", "js"}})

	search.Set(filtered)
	b, err = json.Marshal(search)
	Expect(t, err).ToBe(nil)
	Expect(t, string(b)).ToBe(`{"query":{"filtered":{"query":{"bool":{"should":[{"terms":{"lang":["go","js"]}}]}}}},"sort":[{"hoge":"desc"},"fuga"]}`)

	filtered.Query().Bool().Should(MultiMatch{"k-on", []string{"hobby", "intro"}})
	search.Set(filtered)
	b, err = json.Marshal(search)
	Expect(t, err).ToBe(nil)
	Expect(t, string(b)).ToBe(`{"query":{"filtered":{"query":{"bool":{"should":[{"terms":{"lang":["go","js"]}},{"multi_match":{"fields":["hobby","intro"],"query":"k-on"}}]}}}},"sort":[{"hoge":"desc"},"fuga"]}`)

}
