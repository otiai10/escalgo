package esql

// Term ...
func Term(kv map[string]interface{}) Serializer {
	return createLeaf("term", []Serializer{}, kv)
}
