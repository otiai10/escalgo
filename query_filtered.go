package escalgo

// Filtered provide `{"filtered":{"query":{},"filter":{}}}`.
// Because it can be a child of query (on the root),
type Filtered struct {
	Queries *Queries `json:"query,omitempty"`
	Filters *struct {
		Ands       []Filterable `json:"and,omitempty"`
		Ors        []Filterable `json:"or,omitempty"`
		Filterable `json:",omitempty"`
	} `json:"filter,omitempty"`
}

// Query ...
func (filtered *Filtered) Query() *Filtered {
	if filtered.Queries == nil {
		filtered.Queries = NewQuery()
	}
	return filtered
}

// Bool ...
func (filtered *Filtered) Bool() *Filtered {
	if filtered.Queries.Bool == nil {
		filtered.Queries.Bool = &Bool{}
	}
	return filtered
}

// Should ...
func (filtered *Filtered) Should(queryables ...Queryable) *Filtered {
	filtered.Queries.Bool.Should(queryables...)
	return filtered
}
