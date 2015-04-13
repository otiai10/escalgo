package esql

// Query ...
func Query() Serializer {
	leaf := &Leaf{
		name:     "query",
		Children: map[string]Serializer{},
	}
	return leaf
}

// Query ...
func (d *Leaf) Query(s Serializer) Serializer {
	leaf := &Leaf{
		name: "query",
		Children: map[string]Serializer{
			s.Name(): s,
		},
	}
	d.Children["query"] = leaf
	return d
}
