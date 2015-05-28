package escalgo

// SearchDSL is root node for /_search request body.
type SearchDSL struct {
	Sorts         []Sortable `json:"sort,omitempty"`
	FilteredQuery *Filtered  `json:"query,omitempty"`       //
	PreQuery      Queryable  `json:"query,omitempty"`       // query on root.
	PostFilter    Filterable `json:"post_filter,omitempty"` //
}

// Search ...
func Search(targets ...string) *SearchDSL {
	return &SearchDSL{}
}

// Query ...
func (search *SearchDSL) Query(queryable Queryable) *SearchDSL {
	// omit conflicted "query" field.
	search.FilteredQuery = nil
	// set new "query" field.
	search.PreQuery = queryable
	return search
}

// Filtered ...
func (search *SearchDSL) Filtered(filtered *Filtered) *SearchDSL {
	search.PreQuery = nil
	search.FilteredQuery = filtered
	return search
}

// Sort ...
func (search *SearchDSL) Sort(sortables []Sortable) *SearchDSL {
	search.Sorts = sortables
	return search
}
