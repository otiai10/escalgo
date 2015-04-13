package esql

// Range ...
func Range(serializers ...Serializer) Serializer {
	return createLeaf("range", serializers)
}
