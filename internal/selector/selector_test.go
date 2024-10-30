package selector

import (
	"testing"

	"github.com/MuxN4/gocheerio/internal/dom"
)

func TestSelectorMatching(t *testing.T) {
	testCases := []struct {
		name     string
		html     string
		selector string
		matches  int
	}{
		{
			name:     "Tag Selector",
			html:     "<div><p>First</p><p>Second</p></div>",
			selector: "p",
			matches:  2,
		},
		{
			name:     "ID Selector",
			html:     "<div id='test'><p>Content</p></div>",
			selector: "#test",
			matches:  1,
		},
		{
			name:     "Class Selector",
			html:     "<div class='test'><p class='test'>Content</p></div>",
			selector: ".test",
			matches:  2,
		},
		{
			name:     "Multiple Classes",
			html:     "<div class='test1 test2'><p class='test1'>Content</p></div>",
			selector: ".test1.test2",
			matches:  1,
		},
		{
			name:     "Attribute Selector",
			html:     "<div data-test='value'><p data-test='other'>Content</p></div>",
			selector: "[data-test]",
			matches:  2,
		},
		{
			name:     "Attribute Value Selector",
			html:     "<div data-test='value'><p data-test='other'>Content</p></div>",
			selector: "[data-test='value']",
			matches:  1,
		},
		{
			name:     "Combined Selector",
			html:     "<div class='test' id='unique'><p class='test'>Content</p></div>",
			selector: "div.test#unique",
			matches:  1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			doc, err := dom.NewDocument(tc.html)
			if err != nil {
				t.Fatalf("Failed to parse HTML: %v", err)
			}

			matcher := NewMatcher(tc.selector)

			var matches int
			doc.Root().Each(func(n *dom.Node) bool {
				if matcher.Matches(n) {
					matches++
				}
				return true
			})

			if matches != tc.matches {
				t.Errorf("Expected %d matches, got %d for selector %s",
					tc.matches, matches, tc.selector)
			}
		})
	}
}
