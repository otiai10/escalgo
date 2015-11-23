package escalgo

// Filterable can be a child node under "filter".
// See https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-filters.html
//   - [ ] Compoundable: is possible to include UNSPECIFIED and PLURAL OF filterables.
//     - [ ] and
//     - [ ] bool
//   - [ ] Filterable Element: is SPECIFIED filter element.
//     - [ ] exists
//     - [ ] geo_bounding_box
//     - [ ] geo_distance
//     - [ ] geo_distance_range
//     - [ ] geo_polygon
//     - [ ] geo_shape
//     - [ ] has_child
//     - [ ] has_parent
//   - [ ] Others: TODO :(
//     - [ ] constant_score
//     - [ ] nested
//     - [ ] minimum_should_match
//     - [ ] rewirte
type Filterable interface {
	MarshalFilter() ([]byte, error)
}
