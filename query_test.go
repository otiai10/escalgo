package escalgo

import (
	"testing"

	. "github.com/otiai10/mint"
)

func TestQuery(t *testing.T) {
	query := NewQuery().MatchAll()
	Expect(t, query.Map).ToBe(map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
	})
}
