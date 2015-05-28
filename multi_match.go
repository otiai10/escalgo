package escalgo

import "encoding/json"

// MultiMatch provides `{"multi_match":{"query":"...", "fields":["...","..."]}}`,
// it can be a child node of query and filter.
type MultiMatch struct {
	Query  string   `json:"query,omitempty"`
	Fields []string `json:"fields,omitempty"`
}

// MarshalJSON implements MarshalJSON.
func (mm MultiMatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"multi_match": map[string]interface{}{
			"query":  mm.Query,
			"fields": mm.Fields,
		},
	})
}

// MarshalQuery implements Queryable, alias to MarshalJSON.
func (mm MultiMatch) MarshalQuery() ([]byte, error) {
	return mm.MarshalJSON()
}
