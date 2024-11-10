package selector

import (
	"testing"

	"github.com/MuxN4/gocheerio/internal/dom"
)

// Test helper function to verify matches
func verifyMatches(t *testing.T, doc *dom.Document, matcher *Matcher, expectedNodes []string) bool {
	matched := make(map[string]bool)
	doc.Root().Each(func(n *dom.Node) bool {
		if matcher.Matches(n) {
			matched[n.Node.Data] = true
			t.Logf("Matched node: %v with attributes: %v", n.Node.Data, n.Node.Attr)
		}
		return true
	})

	if len(matched) != len(expectedNodes) {
		t.Errorf("Expected %d matches, got %d matches", len(expectedNodes), len(matched))
		return false
	}

	for _, expected := range expectedNodes {
		if !matched[expected] {
			t.Errorf("Expected to match node %s but didn't", expected)
			return false
		}
	}
	return true
}

func TestSelectorMatching(t *testing.T) {
	testCases := []struct {
		name          string
		html          string
		selector      string
		expectedNodes []string // List of node names that should match
	}{
		{
			name:          "Tag Selector",
			html:          "<div><p>First</p><p>Second</p></div>",
			selector:      "p",
			expectedNodes: []string{"p", "p"},
		},
		{
			name:          "ID Selector",
			html:          "<div id='test'><p>Content</p></div>",
			selector:      "#test",
			expectedNodes: []string{"div"},
		},
		{
			name:          "Class Selector",
			html:          "<div class='test'><p class='test'>Content</p></div>",
			selector:      ".test",
			expectedNodes: []string{"div", "p"},
		},
		{
			name:          "Multiple Classes",
			html:          "<div class='test1 test2'><p class='test1'>Content</p></div>",
			selector:      ".test1.test2",
			expectedNodes: []string{"div"},
		},
		{
			name:          "Attribute Selector",
			html:          "<div data-test='value'><p data-test='other'>Content</p></div>",
			selector:      "[data-test]",
			expectedNodes: []string{"div", "p"},
		},
		{
			name:          "Attribute Value Selector",
			html:          "<div data-test='value'><p data-test='other'>Content</p></div>",
			selector:      "[data-test='value']",
			expectedNodes: []string{"div"},
		},
		{
			name:          "Combined Selector",
			html:          "<div class='test' id='unique'><p class='test'>Content</p></div>",
			selector:      "div.test#unique",
			expectedNodes: []string{"div"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			doc, err := dom.NewDocument(tc.html)
			if err != nil {
				t.Fatalf("Failed to parse HTML: %v", err)
			}

			matcher := NewMatcher(tc.selector)
			verifyMatches(t, doc, matcher, tc.expectedNodes)
		})
	}
}
