package escalgo

import (
	"encoding/json"
	"fmt"
)

// Sortable ...
type Sortable interface {
	Marshalable
}

// Sort ...
type Sort struct {
	Field string
	Desc  bool
}

// MarshalJSON implements Sortable.
func (s *Sort) MarshalJSON() ([]byte, error) {
	if s.Desc {
		return json.Marshal(map[string]string{s.Field: "desc"})
	}
	if s.Field == "_score" {
		return []byte(`"_score"`), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, s.Field)), nil
}
