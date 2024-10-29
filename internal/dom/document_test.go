package dom

import "testing"

func TestDocument(t *testing.T) {
	htmlContent := "<html><body><div>Test</div></body></html>"

	t.Run("Document Creation", func(t *testing.T) {
		doc, err := NewDocument(htmlContent)
		if err != nil {
			t.Fatalf("Failed to create document: %v", err)
		}

		if doc.Root() == nil {
			t.Error("Document root is nil")
		}
	})

	t.Run("Document Rendering", func(t *testing.T) {
		doc, _ := NewDocument(htmlContent)
		rendered, err := doc.Render()
		if err != nil {
			t.Fatalf("Failed to render document: %v", err)
		}

		if rendered == "" {
			t.Error("Rendered document is empty")
		}
	})
}
