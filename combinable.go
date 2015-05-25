package escalgo

import "encoding/json"

type CombinableQuery interface {
	Should(...Queryable) CombinableQuery
}

type Bool struct {
	Shoulds []Queryable `json:"should,omitempty"`
	Musts   []Queryable `json:"must,omitempty"`
}

func (b *Bool) Should(queryables ...Queryable) CombinableQuery {
	b.Shoulds = append(b.Shoulds, queryables...)
	return b
}

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
