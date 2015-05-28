package escalgo

import "encoding/json"

type MultiMatch struct {
	Query  string   `json:"query,omitempty"`
	Fields []string `json:"fields,omitempty"`
}

func (mm MultiMatch) Set(node interface{}) Queryable {
	// do nothing
	return mm
}

func (mm MultiMatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"multi_match": map[string]interface{}{
			"query":  mm.Query,
			"fields": mm.Fields,
		},
	})
}
