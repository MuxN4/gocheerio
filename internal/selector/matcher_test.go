package selector

import (
	"testing"

	"github.com/MuxN4/gocheerio/internal/dom"
	"golang.org/x/net/html"
)

func TestMatchesAttribute(t *testing.T) {
	testsCases := []struct {
		name     string
		node     *dom.Node
		attr     *AttributeSelector
		expected bool
	}{
		{
			name:     "Attribute exists",
			node:     &dom.Node{Node: &html.Node{Type: html.ElementNode, Data: "div", Attr: []html.Attribute{{Key: "id", Val: "test"}}}},
			attr:     &AttributeSelector{Key: "id", Value: "", Operator: ""},
			expected: true,
		},
		{
			name:     "Attribute equals",
			node:     &dom.Node{Node: &html.Node{Type: html.ElementNode, Data: "div", Attr: []html.Attribute{{Key: "id", Val: "test"}}}},
			attr:     &AttributeSelector{Key: "id", Value: "test", Operator: "="},
			expected: true,
		},
		{
			name:     "Attribute not equals",
			node:     &dom.Node{Node: &html.Node{Type: html.ElementNode, Data: "div", Attr: []html.Attribute{{Key: "id", Val: "test"}}}},
			attr:     &AttributeSelector{Key: "id", Value: "not-test", Operator: "="},
			expected: false,
		},
		{
			name:     "Attribute contains word",
			node:     &dom.Node{Node: &html.Node{Type: html.ElementNode, Data: "div", Attr: []html.Attribute{{Key: "class", Val: "foo bar"}}}},
			attr:     &AttributeSelector{Key: "class", Value: "foo", Operator: "~="},
			expected: true,
		},
		{
			name:     "Attribute starts with",
			node:     &dom.Node{Node: &html.Node{Type: html.ElementNode, Data: "div", Attr: []html.Attribute{{Key: "class", Val: "foo bar"}}}},
			attr:     &AttributeSelector{Key: "class", Value: "foo", Operator: "^="},
			expected: true,
		},
		{
			name:     "Attribute ends with",
			node:     &dom.Node{Node: &html.Node{Type: html.ElementNode, Data: "div", Attr: []html.Attribute{{Key: "class", Val: "foo bar"}}}},
			attr:     &AttributeSelector{Key: "class", Value: "bar", Operator: "$="},
			expected: true,
		},
		{
			name:     "Attribute contains",
			node:     &dom.Node{Node: &html.Node{Type: html.ElementNode, Data: "div", Attr: []html.Attribute{{Key: "class", Val: "foo bar"}}}},
			attr:     &AttributeSelector{Key: "class", Value: "oo b", Operator: "*="},
			expected: true,
		},
		{
			name:     "Attribute pipe equals",
			node:     &dom.Node{Node: &html.Node{Type: html.ElementNode, Data: "div", Attr: []html.Attribute{{Key: "lang", Val: "en-US"}}}},
			attr:     &AttributeSelector{Key: "lang", Value: "en", Operator: "|="},
			expected: true,
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.name, func(t *testing.T) {
			matcher := &Matcher{}
			result := matcher.matchesAttribute(tt.node, tt.attr)
			if result != tt.expected {
				t.Errorf("matchesAttribute() = %v, expected %v", result, tt.expected)
			}
		})
	}
}
