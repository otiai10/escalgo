package escalgo

// Queryable can be a query.
type Queryable interface {
	MarshalQuery() ([]byte, error)
}
