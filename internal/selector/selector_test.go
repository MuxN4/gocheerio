package selector

import (
	"testing"

	"github.com/MuxN4/gocheerio/internal/dom"
)

// Test helper function to count unique matches by their data and attributes
func countMatchesByType(t *testing.T, doc *dom.Document, matcher *Matcher, attrKey string) map[string]int {
	matches := make(map[string]int)
	doc.Root().Each(func(n *dom.Node) bool {
		if matcher.Matches(n) {
			t.Logf("Matched node: %v with attributes: %v", n.Node.Data, n.Node.Attr)
			var key string
			if attrKey != "" {
				if val, exists := n.GetAttribute(attrKey); exists {
					key = n.Node.Data + "[" + attrKey + "=" + val + "]"
				}
			} else {
				key = n.Node.Data
			}
			matches[key]++
		}
		return true
	})
	return matches
}

func TestSelectorMatching(t *testing.T) {
	testCases := []struct {
		name     string
		html     string
		selector string
		attrKey  string
		expected map[string]int
	}{
		{
			name:     "Tag Selector",
			html:     "<div><p>First</p><p>Second</p></div>",
			selector: "p",
			expected: map[string]int{"p": 2},
		},
		{
			name:     "ID Selector",
			html:     "<div id='test'><p>Content</p></div>",
			selector: "#test",
			expected: map[string]int{"div": 1},
		},
		{
			name:     "Class Selector",
			html:     "<div class='test'><p class='test'>Content</p></div>",
			selector: ".test",
			expected: map[string]int{"div": 1, "p": 1},
		},
		{
			name:     "Multiple Classes",
			html:     "<div class='test1 test2'><p class='test1'>Content</p></div>",
			selector: ".test1.test2",
			expected: map[string]int{"div": 1},
		},
		{
			name:     "Attribute Selector",
			html:     "<div data-test='value'><p data-test='other'>Content</p></div>",
			selector: "[data-test]",
			attrKey:  "data-test",
			expected: map[string]int{
				"div[data-test=value]": 1,
				"p[data-test=other]":   1,
			},
		},
		{
			name:     "Attribute Value Selector",
			html:     "<div data-test='value'><p data-test='other'>Content</p></div>",
			selector: "[data-test='value']",
			attrKey:  "data-test",
			expected: map[string]int{
				"div[data-test=value]": 1,
			},
		},
		{
			name:     "Combined Selector",
			html:     "<div class='test' id='unique'><p class='test'>Content</p></div>",
			selector: "div.test#unique",
			expected: map[string]int{"div": 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			doc, err := dom.NewDocument(tc.html)
			if err != nil {
				t.Fatalf("Failed to parse HTML: %v", err)
			}

			matcher := NewMatcher(tc.selector)
			matches := countMatchesByType(t, doc, matcher, tc.attrKey)

			// Compare results
			if len(matches) != len(tc.expected) {
				t.Errorf("Expected %d different types of matches, got %d", len(tc.expected), len(matches))
			}

			for key, expectedCount := range tc.expected {
				if gotCount := matches[key]; gotCount != expectedCount {
					t.Errorf("For %s: expected %d matches, got %d", key, expectedCount, gotCount)
				}
			}
		})
	}
}
