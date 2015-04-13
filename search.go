package esql

// Search ...
func Search(indexType string) Serializer {
	return &Leaf{
		Children: map[string]Serializer{},
	}
}
