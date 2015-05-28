package escalgo

// MatchAll provides `{"match_all":{}}`, it only can be a query's child node.
type MatchAll struct {
}

// MarshalJSON implements json.Marshaler.
func (ma *MatchAll) MarshalJSON() ([]byte, error) {
	return []byte(`{"match_all":{}}`), nil
}

// MarshalQuery implements Queryable, alias to MarshalJSON.
func (ma *MatchAll) MarshalQuery() ([]byte, error) {
	return ma.MarshalJSON()
}
