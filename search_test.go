package escalgo

import (
	"testing"

	. "github.com/otiai10/mint"
)

func SearchRootTest(t *testing.T) {
	Expect(t, true).ToBe(true)

	search := Search()

	query := NewQuery().MatchAll()
	filter := NewFilter()

	search.Query.Filtered.Query = query
	search.Query.Filtered.Filter = filter

	m, err := search.Marshal()

    Expect(t, m).TypeOf("[]byte")
    Expect(t, err).ToBe(nil)
}
