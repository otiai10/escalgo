package sdsl

// Filterable is an interface for Filter.
// Any types implementing `Filterable` are able to be Filter.
type Filterable interface{}

// Filterables is just a slice for Filterable.
type Filterables []Filterable
