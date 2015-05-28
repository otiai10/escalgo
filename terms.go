package escalgo

import "encoding/json"

type Terms struct {
	Field string
	Words []string
}

func (terms Terms) Set(node interface{}) Queryable {
	// do nothing
	return terms
}

func (terms Terms) MarshalJSON() ([]byte, error) {
	t := map[string]interface{}{}
	t[terms.Field] = terms.Words
	m := map[string]interface{}{
		"terms": t,
	}
	return json.Marshal(m)
}
