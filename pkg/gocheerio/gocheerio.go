package gocheerio

// Document represents an HTML document
type Document interface {
    // Find returns a Selection object containing descendants of the current document
    // that match the provided selector
    Find(selector string) Selection

    // Html returns the HTML contents of the first element in the selection
    Html() (string, error)

    // Text returns the combined text contents of the document
    Text() string
}

// Selection represents a collection of nodes matched by selectors
type Selection interface {
    // Find returns a new Selection containing descendants that match the selector
    Find(selector string) Selection

    // Html returns the HTML contents of the first element in the selection
    Html() (string, error)

    // Text returns the combined text contents of all matched elements
    Text() string

    // Attr returns the value of the specified attribute for the first element in the Selection
    Attr(name string) (string, bool)
}

// Load creates a new Document from the provided HTML string
func Load(html string) Document {
    // Implementation will be added as I proceed
    panic("Not implemented")
}