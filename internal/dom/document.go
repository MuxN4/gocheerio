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
	doc.root = NewNode(node, doc)
	return doc, nil
}

// Root returns the document's root node.
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
