package escalgo

type SearchRoot struct {
	Query struct {
		Filtered struct {
			Query  *Query  `json:"query"`
			Filter *Filter `json:"filter"`
		} `json:"filtered"`
	} `json:"query"`
	Sort struct{} `json:"sort"`
}

func Search() *SearchRoot {
	return &SearchRoot{}
}

func (root *SearchRoot) Marshal() ([]byte, error) {
    return []byte{}, nil
}
