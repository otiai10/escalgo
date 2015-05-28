package escalgo

// SearchDSL is root node for /_search request body.
type SearchDSL struct {
	Query      Queryable  `json:"query,omitempty"`
	Sort       []Sortable `json:"sort,omitempty"`
	PostFilter Filterable `json:"post_filter,omitempty"`
}

// Search ...
func Search(defs ...string) *SearchDSL {
	return &SearchDSL{}
}

// Set is setter for field for root.
func (search *SearchDSL) Set(node interface{}) *SearchDSL {
	switch n := node.(type) {
	case []Sortable:
		search.Sort = n
	case *Filtered:
		search.Query = NewQuery()
		search.Query.Set(n)
	}
	return search
}
