package escalgo

// Filterable can be a child node under "filter".
// See https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-filters.html
//   - [ ] Compoundable: is possible to include UNSPECIFIED and PLURAL OF filterables.
//     - [ ] and
//     - [ ] bool
//   - [ ] Filterable Element: is SPECIFIED filter element.
//   - [ ] Others: TODO :(
//     - [ ] constant_score
//     - [ ] nested
//     - [ ] minimum_should_match
//     - [ ] rewirte
type Filterable interface {
	MarshalFilter() ([]byte, error)
}
