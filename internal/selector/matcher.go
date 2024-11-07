package selector

import (
	"strings"

	"github.com/MuxN4/gocheerio/internal/dom"
	"golang.org/x/net/html"
)

type Matcher struct {
	selectors []*Selector
}

func NewMatcher(selector string) *Matcher {
	parser := NewParser(selector)
	return &Matcher{
		selectors: parser.Parse(),
	}
}

func (m *Matcher) Matches(node *dom.Node) bool {
	// For compound selectors, all parts must match
	for _, sel := range m.selectors {
		if !m.matchesSelector(node, sel) {
			return false
		}
	}
	return true
}

func (m *Matcher) matchesSelector(node *dom.Node, sel *Selector) bool {
	// Only match element nodes
	if node.Node.Type != html.ElementNode {
		return false
	}

	// Check tag
	if sel.Tag != "" && sel.Tag != node.Node.Data {
		return false
	}

	// Check ID
	if sel.ID != "" {
		if id, exists := node.GetAttribute("id"); !exists || id != sel.ID {
			return false
		}
	}

	// Check classes
	if len(sel.Classes) > 0 {
		nodeClass, exists := node.GetAttribute("class")
		if !exists {
			return false
		}

		classes := strings.Fields(nodeClass)
		classMap := make(map[string]bool)
		for _, class := range classes {
			classMap[class] = true
		}

		for _, class := range sel.Classes {
			if !classMap[class] {
				return false
			}
		}
	}

	// Check attribute
	if sel.Attribute != nil {
		if !m.matchesAttribute(node, sel.Attribute) {
			return false
		}
	}

	return true
}

func (m *Matcher) matchesAttribute(node *dom.Node, attr *AttributeSelector) bool {
	// Only match element nodes
	if node.Node.Type != html.ElementNode {
		return false
	}

	// Gets the attribute value
	value, exists := node.GetAttribute(attr.Key)
	if !exists {
		return false
	}

	// If only checking for attribute existence
	if attr.Value == "" && attr.Operator == "" {
		return true
	}

	// Clean the attribute value we're looking for
	expectedValue := strings.Trim(attr.Value, "'\"")

	// If no operator is specified but value given, use exact match
	if attr.Operator == "" {
		return value == expectedValue
	}

	switch attr.Operator {
	case "=":
		return value == expectedValue
	case "~=":
		return containsWord(value, expectedValue)
	case "|=":
		return value == expectedValue || strings.HasPrefix(value, expectedValue+"-")
	case "^=":
		return strings.HasPrefix(value, expectedValue)
	case "$=":
		return strings.HasSuffix(value, expectedValue)
	case "*=":
		return strings.Contains(value, expectedValue)
	}

	return false
}

func containsWord(s, word string) bool {
	words := strings.Fields(s)
	for _, w := range words {
		if w == word {
			return true
		}
	}
	return false
}
