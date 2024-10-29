package dom

import (
	"testing"

	"golang.org/x/net/html"
)

func TestSelection(t *testing.T) {
	htmlContent := `
        <div>
            <span>First</span>
            <span>Second</span>
            <span>Third</span>
        </div>
    `

	doc, _ := NewDocument(htmlContent)

	t.Run("Selection Methods", func(t *testing.T) {
		// Find all spans
		spans := doc.Root().FindNodes(func(n *Node) bool {
			return n.Node.Type == html.ElementNode && n.Node.Data == "span"
		})

		if len(spans.nodes) != 3 {
			t.Errorf("Expected 3 spans, got %d", len(spans.nodes))
		}

		first := spans.First()
		if first.nodes[0].Node.FirstChild.Data != "First" {
			t.Error("First() failed")
		}

		last := spans.Last()
		if last.nodes[0].Node.FirstChild.Data != "Third" {
			t.Error("Last() failed")
		}

		second := spans.Eq(1)
		if second.nodes[0].Node.FirstChild.Data != "Second" {
			t.Error("Eq(1) failed")
		}
	})
}
