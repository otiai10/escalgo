package escalgo

// Queryable can be a child node under "query".
// ** https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl.html
// ** https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-queries.html
// - [ ] Compoundable: is possible to include UNSPECIFIED and PLURAL OF queryables.
//   - [ ] bool
//   - [ ] boosting
//   - [ ] dis_max
//   - [ ] filtered
//   - [ ] has_child
//   - [ ] has_parent
//   - [ ] indices
//   - [ ] span_first
//   - [ ] span_multi
//   - [ ] span_near
//   - [ ] span_not
//   - [ ] span_or
//   - [ ] top_children
// - [ ] Queryable Element: is SPECIFIED query element.
//   - [ ] match
//   - [ ] common
//   - [ ] fuzzy_like_this
//   - [ ] fuzzy_like_this_field
//   - [ ] function_score
//   - [ ] fuzzy
//   - [ ] geo_shape
//   - [ ] ids
//   - [ ] match_all
//   - [ ] more_like_this
//   - [ ] prefix
//   - [ ] query_string
//   - [ ] simple_query_string
//   - [ ] range
//   - [ ] regexp
//   - [ ] span_term
//   - [ ] term
//   - [ ] terms
//   - [ ] wildcard
//   - [ ] template
// - [ ] Others: TODO :(
//   - [ ] constant_score
//   - [ ] nested
//   - [ ] minimum_should_match
//   - [ ] rewirte
type Queryable interface {
	Set(interface{}) Queryable
}
