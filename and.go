package esql

// And ...
func And(serializers ...Serializer) Serializer {
	return createLeaf("and", serializers)
}
