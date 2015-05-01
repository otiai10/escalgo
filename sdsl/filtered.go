package sdsl

// Filtered represents "filtered" filed on the top level in search params.
type Filtered struct {
	Query      Query
	Filter     Filter
	PostFilter Filter
}
