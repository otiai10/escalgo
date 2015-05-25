package escalgo

import "encoding/json"

// Marshalable ...
type Marshalable interface {
	json.Marshaler
}
