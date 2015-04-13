package esql

// Serializer ...
type Serializer interface {
	Serialize() map[string]interface{}
	Query(Serializer) Serializer
	Name() string
	String(...bool) string
}
