package dom

import (
	"bytes"

	"golang.org/x/net/html"
)

// Document represents an HTML document with a root node
type Document struct {
	root *Node
}

// NewDocument creates a new Document instance from an HTML string
func NewDocument(htmlContent string) (*Document, error) {
	node, err := html.Parse(bytes.NewReader([]byte(htmlContent)))
	if err != nil {
		return nil, err
	}

	doc := &Document{}
	doc.root = findFirstElement(node)
	return doc, nil
}

// findFirstElement, recursively finds the first element node
func findFirstElement(node *html.Node) *Node {
	if node.Type == html.ElementNode {
		return NewNode(node, nil)
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		if result := findFirstElement(child); result != nil {
			return result
		}
	}

	return nil
}

// Root returns the document's root node for querying
func (d *Document) Root() *Node {
	return d.root
}

// Render returns the HTML representation of the document
func (d *Document) Render() (string, error) {
	var buf bytes.Buffer
	err := html.Render(&buf, d.root.Node)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
