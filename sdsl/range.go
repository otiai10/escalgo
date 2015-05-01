package sdsl

// Range represents "range" field of any places in search params.
// `Range` can be a child of "filter", so `Range` implements `Filterable`.
type Range struct {
	Filed string
	Gt    interface{}
	Gte   interface{}
	Lt    interface{}
	Lte   interface{}
}
