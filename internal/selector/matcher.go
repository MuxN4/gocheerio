package selector

import (
	"strings"

	"github.com/MuxN4/gocheerio/internal/dom"
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
	// Check tag
	if sel.Tag != "" && sel.Tag != node.Data {
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
	value, exists := node.GetAttribute(attr.Key)
	if !exists {
		return false
	}

	if attr.Value == "" {
		return true // Just checking existence
	}

	switch attr.Operator {
	case "=":
		return value == attr.Value
	case "~=":
		return containsWord(value, attr.Value)
	case "|=":
		return value == attr.Value || strings.HasPrefix(value, attr.Value+"-")
	case "^=":
		return strings.HasPrefix(value, attr.Value)
	case "$=":
		return strings.HasSuffix(value, attr.Value)
	case "*=":
		return strings.Contains(value, attr.Value)
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
