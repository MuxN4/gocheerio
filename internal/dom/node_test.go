// node_test.go
package dom

import (
	"testing"

	"golang.org/x/net/html"
)

func TestNode(t *testing.T) {
	htmlContent := `
        <html>
            <body>
                <div id="parent">
                    <span class="child">Child 1</span>
                    <span class="child">Child 2</span>
                </div>
            </body>
        </html>
    `

	doc, err := NewDocument(htmlContent)
	if err != nil {
		t.Fatalf("Failed to create document: %v", err)
	}

	t.Run("Node Navigation", func(t *testing.T) {
		// Find the div node
		var divNode *Node
		doc.Root().Each(func(n *Node) bool {
			if n.Node.Type == html.ElementNode && n.Node.Data == "div" {
				divNode = n
				return false
			}
			return true
		})

		if divNode == nil {
			t.Fatal("Failed to find div node")
		}

		// Test parent
		parent := divNode.Parent()
		if parent == nil || parent.Node.Data != "body" {
			t.Errorf("Expected parent to be body, got %v", parent)
		}

		// Test children
		firstChild := divNode.FirstChild()
		if firstChild == nil {
			t.Error("First child is nil")
		} else if firstChild.Node.Data != "span" {
			t.Errorf("Expected first child to be span, got %s", firstChild.Node.Data)
		}

		lastChild := divNode.LastChild()
		if lastChild == nil {
			t.Error("Last child is nil")
		} else if lastChild.Node.Data != "span" {
			t.Errorf("Expected last child to be span, got %s", lastChild.Node.Data)
		}

		// Additional test to verify span content
		if firstChild != nil {
			firstText := firstChild.Node.FirstChild
			if firstText == nil || firstText.Data != "Child 1" {
				t.Errorf("Expected first span content to be 'Child 1', got %v", firstText)
			}
		}

		if lastChild != nil {
			lastText := lastChild.Node.FirstChild
			if lastText == nil || lastText.Data != "Child 2" {
				t.Errorf("Expected last span content to be 'Child 2', got %v", lastText)
			}
		}
	})
}
