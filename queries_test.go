package escalgo

import (
	"encoding/json"
	"testing"

	. "github.com/otiai10/mint"
)

func TestQueries(t *testing.T) {
	Expect(t, true).ToBe(true)

	queries := NewQuery().MatchAll()

	b, err := json.Marshal(queries)
	Expect(t, err).ToBe(nil)
	Expect(t, string(b)).ToBe(`{}`)
}
