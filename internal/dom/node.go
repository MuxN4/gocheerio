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

func (n *Node) FirstChild() *Node {
	if n.Node.FirstChild == nil {
		return nil
	}
	return NewNode(n.Node.FirstChild, n.document)
}

func (n *Node) LastChild() *Node {
	if n.Node.LastChild == nil {
		return nil
	}
	return NewNode(n.Node.LastChild, n.document)
}

func (n *Node) NextSibling() *Node {
	if n.Node.NextSibling == nil {
		return nil
	}
	return NewNode(n.Node.NextSibling, n.document)
}

func (n *Node) PrevSibling() *Node {
	if n.Node.PrevSibling == nil {
		return nil
	}
	return NewNode(n.Node.PrevSibling, n.document)
}