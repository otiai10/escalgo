package escalgo

import "encoding/json"

// Terms provides `{"terms":{"your_field":["...","..."]}}`,
// it can be a child node of query and filter.
type Terms struct {
	Field string
	Words []string
}

// MarshalJSON to implement json.Marshaler.
func (terms Terms) MarshalJSON() ([]byte, error) {
	t := map[string]interface{}{}
	t[terms.Field] = terms.Words
	m := map[string]interface{}{
		"terms": t,
	}
	return json.Marshal(m)
}

// MarshalQuery to implement Queryable, alias to MarshalJSON.
func (terms Terms) MarshalQuery() ([]byte, error) {
	return terms.MarshalJSON()
}
