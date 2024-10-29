package dom

import "golang.org/x/net/html"

// Node wraps the html.Node type with additional functionality
type Node struct {
	*html.Node
	document *Document
}

// NewNode creates a new Node instance
func NewNode(node *html.Node, doc *Document) *Node {
	return &Node{
		Node:     node,
		document: doc,
	}
}

// Parent returns the parent node
func (n *Node) Parent() *Node {
	if n.Node.Parent == nil {
		return nil
	}
	return NewNode(n.Node.Parent, n.document)
}

// FirstChild returns the first child element node, skipping non-element nodes
func (n *Node) FirstChild() *Node {
	current := n.Node.FirstChild
	for current != nil {
		if current.Type == html.ElementNode {
			return NewNode(current, n.document)
		}
		current = current.NextSibling
	}
	return nil
}

// LastChild returns the last child element node, skipping non-element nodes
func (n *Node) LastChild() *Node {
	current := n.Node.LastChild
	for current != nil {
		if current.Type == html.ElementNode {
			return NewNode(current, n.document)
		}
		current = current.PrevSibling
	}
	return nil
}

func (n *Node) NextSibling() *Node {
	current := n.Node.NextSibling
	for current != nil {
		if current.Type == html.ElementNode {
			return NewNode(current, n.document)
		}
		current = current.NextSibling
	}
	return nil
}

func (n *Node) PrevSibling() *Node {
	current := n.Node.PrevSibling
	for current != nil {
		if current.Type == html.ElementNode {
			return NewNode(current, n.document)
		}
		current = current.PrevSibling
	}
	return nil
}

// Each traverses nodes, calling the callback, stops if callback returns false
func (n *Node) Each(callback func(*Node) bool) bool {
	if !callback(n) {
		return false
	}

	for child := n.Node.FirstChild; child != nil; child = child.NextSibling {
		if child.Type == html.ElementNode {
			childNode := NewNode(child, n.document)
			if !childNode.Each(callback) {
				return false
			}
		}
	}
	return true
}

// FindNodes returns all descendant nodes matching the specified criteria
func (n *Node) FindNodes(matcher func(*Node) bool) *Selection {
	var matches []*Node
	n.Each(func(node *Node) bool {
		if matcher(node) {
			matches = append(matches, node)
		}
		return true
	})
	return NewSelection(matches, n.document)
}
