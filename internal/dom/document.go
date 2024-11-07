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

// Root returns the document's root node for querying
func (d *Document) Root() *Node {
	// Find the body element
	var body *Node
	d.root.Each(func(n *Node) bool {
		if n.Node.Type == html.ElementNode && n.Node.Data == "body" {
			body = n
			return false
		}
		return true
	})

	// If body is found, return its first element child
	if body != nil {
		return NewNode(body.Node.FirstChild, d)
	}

	// Fallback to document root if no body found
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
