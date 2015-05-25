package escalgo

// Queryable can be a query.
type Queryable interface {
	Set(interface{}) Queryable
}
