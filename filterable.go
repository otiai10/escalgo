package escalgo

// Filterable can be a filter.
type Filterable interface {
	MashalFilter() ([]byte, error)
}
