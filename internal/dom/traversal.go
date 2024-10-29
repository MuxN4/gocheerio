package dom

// FindNodes returns all descendant nodes matching the specified criteria
func (s *Selection) FindNodes(matcher func(*Node) bool) *Selection {
	var matches []*Node
	for _, node := range s.nodes {
		matches = append(matches, findRecursive(node, matcher)...)
	}
	return NewSelection(matches, s.doc)
}

func findRecursive(node *Node, matcher func(*Node) bool) []*Node {
	var matches []*Node

	if matcher(node) {
		matches = append(matches, node)
	}

	for child := node.FirstChild(); child != nil; child = child.NextSibling() {
		matches = append(matches, findRecursive(child, matcher)...)
	}

	return matches
}
