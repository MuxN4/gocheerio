package dom

// Selection represents a collection of nodes
type Selection struct {
	nodes []*Node
	doc   *Document
}

// NewSelection creates a new Selection instance
func NewSelection(nodes []*Node, doc *Document) *Selection {
	return &Selection{
		nodes: nodes,
		doc:   doc,
	}
}

// First returns the first element in the selection or an empty Selection if none exists.
func (s *Selection) First() *Selection {
	if len(s.nodes) == 0 {
		return NewSelection(nil, s.doc)
	}
	return NewSelection([]*Node{s.nodes[0]}, s.doc)
}

func (s *Selection) Last() *Selection {
	if len(s.nodes) == 0 {
		return NewSelection(nil, s.doc)
	}
	return NewSelection([]*Node{s.nodes[len(s.nodes)-1]}, s.doc)
}

// Eq returns the element at the specified index or an empty Selection if out of bounds
func (s *Selection) Eq(index int) *Selection {
	if index < 0 || index >= len(s.nodes) {
		return NewSelection(nil, s.doc)
	}
	return NewSelection([]*Node{s.nodes[index]}, s.doc)
}

// Each iterates over the selection and executes the callback for each node
func (s *Selection) Each(callback func(int, *Node)) {
	for i, node := range s.nodes {
		callback(i, node)
	}
}

// Nodes returns the underlying node slice
func (s *Selection) Nodes() []*Node {
	return s.nodes
}
