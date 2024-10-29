package gocheerio

import "github.com/MuxN4/gocheerio/internal/dom"

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

	// Attr returns the value of the specified attribute for the first element in the Selection
	Attr(name string) (string, bool)
}

type document struct {
	doc *dom.Document
}

type selection struct {
	sel *dom.Selection
}

// Load creates a new Document from HTML content
func Load(html string) Document {
	doc, err := dom.NewDocument(html)
	if err != nil {
		panic(err) // For now, I'll panic on error. Later error handling can be improved
	}
	return &document{doc: doc}
}

func (d *document) Find(selector string) Selection {
	// TODO: Implement selector matching
	// For now, return empty selection
	return &selection{sel: dom.NewSelection(nil, d.doc)}
}

func (d *document) Html() (string, error) {
	return d.doc.Render()
}

func (d *document) Text() string {
	// TODO: Implement text extraction
	return ""
}

func (s *selection) Find(selector string) Selection {
	// TODO: Implement selector matching
	return s
}

func (s *selection) Html() (string, error) {
	// TODO: Implement HTML rendering for selection
	return "", nil
}

func (s *selection) Text() string {
	// TODO: Implement text extraction
	return ""
}

func (s *selection) Attr(name string) (string, bool) {
	// TODO: Implement attribute access
	return "", false
}
