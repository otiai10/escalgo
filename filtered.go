package escalgo

// Filtered ...
type Filtered struct {
	Queries *Queries `json:"query,omitempty"`
	/*
		Queries *struct {
			Bool *struct {
				Shoulds []Queryable `json:"should,omitempty"`
				Musts   []Queryable `json:"must,omitempty"`
			} `json:"bool,omitempty"`
			Queryable `json:",omitempty"`
		} `json:"query,omitempty"`
	*/
	Filters *struct {
		Ands       []Filterable `json:"and,omitempty"`
		Ors        []Filterable `json:"or,omitempty"`
		Filterable `json:",omitempty"`
	} `json:"filter,omitempty"`
}

func (filtered *Filtered) Query() *Filtered {
	if filtered.Queries == nil {
		filtered.Queries = NewQuery()
	}
	return filtered
}

func (filtered *Filtered) Bool() *Filtered {
	if filtered.Queries.Bool == nil {
		filtered.Queries.Bool = &Bool{}
	}
	return filtered
}

func (filtered *Filtered) Should(queryables ...Queryable) *Filtered {
	filtered.Queries.Bool.Should(queryables...)
	return filtered
}
