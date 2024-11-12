package selector

import (
	"fmt"
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
	if node.Node.Type != html.ElementNode {
		return false
	}

	// For compound selectors, all parts must match
	for _, sel := range m.selectors {
		if !m.matchesSelector(node, sel) {
			return false
		}
	}
	return true
}

func (m *Matcher) matchesSelector(node *dom.Node, sel *Selector) bool {
	// !If there's an attribute selector, check it first
	if sel.Attribute != nil {
		if !m.matchesAttribute(node, sel.Attribute) {
			return false
		}
	}

	// *Check tag if specified
	if sel.Tag != "" && sel.Tag != node.Node.Data {
		return false
	}

	// *Check ID if specified
	if sel.ID != "" {
		if id, exists := node.GetAttribute("id"); !exists || id != sel.ID {
			return false
		}
	}

	// *Check classes if specified
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

	return true
}

func (m *Matcher) matchesAttribute(node *dom.Node, attr *AttributeSelector) bool {
	value, exists := node.GetAttribute(attr.Key)
	fmt.Printf("Node: %s, Key: %s, Value: %s, Exists: %v\n", node.Node.Data, attr.Key, value, exists)
	fmt.Printf("Selector Value: %s, Operator: %s\n", attr.Value, attr.Operator)

	// If the attribute doesn't exist, no match
	if !exists {
		return false
	}

	// !If we're only checking for attribute existence
	if attr.Value == "" && attr.Operator == "" {
		return true
	}

	expectedValue := strings.Trim(attr.Value, "'\"")
	fmt.Printf("Expected Value (after trim): %s\n", expectedValue)

	// For value matching (both with = operator or implicit)
	if attr.Operator == "" || attr.Operator == "=" {
		return value == expectedValue
	}

	switch attr.Operator {
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
