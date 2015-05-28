package escalgo

import "encoding/json"

// CompoundableQuery ...
type CompoundableQuery interface {
	Queryable
}

// Bool represents {"bool":{...}}.
// ** https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-bool-query.html
type Bool struct {
	Shoulds  []Queryable `json:"should,omitempty"`
	Musts    []Queryable `json:"must,omitempty"`
	MustNots []Queryable `json:"must_not,omitempty"`
}

// Should of Bool add element for "should".
func (b *Bool) Should(queryables ...Queryable) *Bool {
	b.Shoulds = append(b.Shoulds, queryables...)
	return b
}

// MarshalJSON to implement json.Marshaler.
func (b *Bool) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{}
	if len(b.Shoulds) != 0 {
		m["should"] = b.Shoulds
	}
	if len(b.Musts) != 0 {
		m["must"] = b.Musts
	}
	return json.Marshal(m)
}
