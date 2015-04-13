package esql

import "encoding/json"

// Leaf ...
type Leaf struct {
	Children map[string]Serializer
	name     string
	values   map[string]interface{}
}

// Serialize ...
func (d *Leaf) Serialize() map[string]interface{} {
	if len(d.values) != 0 {
		return d.values
	}
	children := map[string]interface{}{}
	for name, serializer := range d.Children {
		children[name] = serializer.Serialize()
	}
	return children
}

// Name ...
func (d *Leaf) Name() string {
	return d.name
}

func (d *Leaf) String(pretty ...bool) string {
	b, _ := json.Marshal(d.Serialize())
	return string(b)
}

func createLeaf(name string, children []Serializer, values ...map[string]interface{}) *Leaf {
	leaf := &Leaf{
		name:     name,
		Children: map[string]Serializer{},
		values:   map[string]interface{}{},
	}
	for _, child := range children {
		leaf.Children[child.Name()] = child
	}
	if len(values) != 0 {
		leaf.values = values[0]
	}
	return leaf
}
