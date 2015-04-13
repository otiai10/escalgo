package esql

// Filtered ...
func Filtered(serializers ...Serializer) Serializer {
	return createLeaf("filtered", serializers)
}
