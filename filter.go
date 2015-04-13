package esql

// Filter ...
func Filter(serializers ...Serializer) Serializer {
	return createLeaf("filter", serializers)
}
