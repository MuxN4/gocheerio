package gocheerio

import (
	"bytes"
	"strings"

	"github.com/MuxN4/gocheerio/internal/dom"
	"github.com/MuxN4/gocheerio/internal/selector"
	"golang.org/x/net/html"
)

// Document represents an HTML document
type Document interface {
	// Find returns a Selection matching the given selector
	Find(selector string) Selection

	// Html returns the HTML content of the document
	Html() (string, error)

	// Text returns the text content of the document
	Text() string
}

// Selection represents a set of matched nodes
type Selection interface {
	// Find returns a new Selection filtered by the given selector
	Find(selector string) Selection

	// Html returns the HTML contents of the first element in the selection
	Html() (string, error)

	// Text returns the combined text contents of all matched elements
	Text() string

	// Attr returns the value of the specified attribute for the first element
	Attr(name string) (string, bool)

	// Each runs the given function on each element in the selection
	Each(func(int, Selection))

	// Length returns the number of elements in the selection
	Length() int
}

type document struct {
	doc *dom.Document
}

type selection struct {
	sel *dom.Selection
	doc *dom.Document // *Added document reference
}

// Load creates a new Document from HTML content
func Load(html string) (Document, error) {
	doc, err := dom.NewDocument(html)
	if err != nil {
		return nil, err
	}
	return &document{doc: doc}, nil
}

func (d *document) Find(selectorStr string) Selection {
	if selectorStr == "" {
		return &selection{sel: dom.NewSelection(nil, d.doc), doc: d.doc}
	}

	matcher := selector.NewMatcher(selectorStr)
	matches := make([]*dom.Node, 0)

	// *Use the DOM's traversal function with the matcher
	d.doc.Root().Each(func(n *dom.Node) bool {
		if matcher.Matches(n) {
			matches = append(matches, n)
		}
		return true
	})

	return &selection{sel: dom.NewSelection(matches, d.doc), doc: d.doc}
}

func (d *document) Html() (string, error) {
	return d.doc.Render()
}

func (d *document) Text() string {
	if d.doc.Root() == nil {
		return ""
	}
	var text string
	d.doc.Root().Each(func(n *dom.Node) bool {
		if n.Node.Type == html.TextNode {
			text += n.Node.Data
		}
		return true
	})
	return text
}

func (s *selection) Find(selectorStr string) Selection {
	if selectorStr == "" {
		return &selection{sel: dom.NewSelection(nil, s.doc), doc: s.doc}
	}

	matcher := selector.NewMatcher(selectorStr)
	matches := make([]*dom.Node, 0)

	for _, node := range s.sel.Nodes() {
		node.Each(func(n *dom.Node) bool {
			if matcher.Matches(n) {
				matches = append(matches, n)
			}
			return true
		})
	}

	return &selection{sel: dom.NewSelection(matches, s.doc), doc: s.doc}
}

func (s *selection) Html() (string, error) {
	nodes := s.sel.Nodes()
	if len(nodes) == 0 {
		return "", nil
	}

	// !Use bytes.Buffer to render HTML
	var buf bytes.Buffer
	err := html.Render(&buf, nodes[0].Node)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (s *selection) Text() string {
	var texts []string
	s.sel.Each(func(i int, n *dom.Node) {
		if n.Node.Type == html.TextNode {
			texts = append(texts, n.Node.Data)
		}
	})
	return strings.Join(texts, " ")
}

func (s *selection) Attr(name string) (string, bool) {
	nodes := s.sel.Nodes()
	if len(nodes) == 0 {
		return "", false
	}
	return nodes[0].GetAttribute(name)
}

func (s *selection) Each(f func(int, Selection)) {
	s.sel.Each(func(i int, n *dom.Node) {
		f(i, &selection{
			sel: dom.NewSelection([]*dom.Node{n}, s.doc),
			doc: s.doc,
		})
	})
}

func (s *selection) Length() int {
	return len(s.sel.Nodes())
}
