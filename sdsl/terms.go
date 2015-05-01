package sdsl

// Terms represents "term" field of any places in search params.
// `Terms` can be a child of "filter", so `Terms` implements `Filterable`.
type Terms struct {
	Filed     string
	Values    []interface{}
	Execution string
}
