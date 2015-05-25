package escalgo

type MatchAll struct {
}

func (ma *MatchAll) MarshalJSON() ([]byte, error) {
	return []byte(`{"match_all":{}}`), nil
}

func (ma *MatchAll) Set(node interface{}) Queryable {
	return ma
}
