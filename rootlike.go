package escalgo

// RootLike can be a root-like node, you know `{"query":{},"filter":{},"sort":[]}`.
// For example, the root node of `/_search` body,
// or `filtered` node in query node.
type RootLike interface {
	Set(interface{}) RootLike
}
