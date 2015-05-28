package escalgo

import "encoding/json"

// Queries represents a node for `"query": {}`.
// MustAll
type Queries struct {
	isMatchAll bool      `json:"-"`
	Only       Queryable `json:",omitempty"` // embed Queryable
	Bool       *Bool     `json:"bool,omitempty"`
	Filtered   *Filtered `json:"filtered,omitempty"`
}

// NewQuery initalize "query" node.
func NewQuery() *Queries {
	return &Queries{}
}

// Must ... (or `And` is better?)
func (queries *Queries) Must(queryables ...Queryable) *Queries {
	return queries
}

// Should ... (or `Or` is better?)
func (queries *Queries) Should(queryables ...Queryable) *Queries {
	return queries
}

// MatchAll ...
func (queries *Queries) MatchAll() *Queries {
	queries.isMatchAll = true
	return queries
}

// MarshalJSON to implements json.Marhsaler.
func (queries *Queries) MarshalJSON() ([]byte, error) {
	if queries.Filtered != nil {
		return json.Marshal(map[string]interface{}{
			"filtered": queries.Filtered,
		})
	}
	if queries.Bool != nil {
		return json.Marshal(map[string]interface{}{
			"bool": queries.Bool,
		})
	}
	return json.Marshal(map[string]interface{}{})
}

func (queries *Queries) Set(node interface{}) Queryable {
	switch n := node.(type) {
	case *Filtered:
		queries.Filtered = n
	}
	return queries
}
