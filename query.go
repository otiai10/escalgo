package escalgo

type Query struct {
	Map map[string]interface{}
}

func NewQuery() *Query {
	return &Query{
		Map: map[string]interface{}{
			"query": map[string]interface{}{},
		},
	}
}

func (query *Query) MatchAll() *Query {
	query.Map["query"] = map[string]interface{}{
		"match_all": map[string]interface{}{},
	}
	return query
}
