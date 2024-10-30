package selector

import "strings"

type SelectorType int

const (
	SelectorTypeTag SelectorType = iota
	SelectorTypeID
	SelectorTypeClass
	SelectorTypeUniversal
	SelectorTypeCompound
)

type Selector struct {
	Type      SelectorType
	Value     string
	Classes   []string
	ID        string
	Tag       string
	Attribute *AttributeSelector
}

type AttributeSelector struct {
	Key      string
	Value    string
	Operator string // =, ~=, |=, ^=, $=, *=
}

type Parser struct {
	tokenizer *Tokenizer
}

func NewParser(selector string) *Parser {
	return &Parser{
		tokenizer: NewTokenizer(selector),
	}
}

func (p *Parser) Parse() []*Selector {
	var selectors []*Selector
	current := &Selector{Type: SelectorTypeCompound}

	for token := p.tokenizer.nextToken(); token != nil; token = p.tokenizer.nextToken() {
		switch token.Type {
		case TokenTypeTag:
			current.Tag = token.Value
		case TokenTypeID:
			current.ID = token.Value
		case TokenTypeClass:
			current.Classes = append(current.Classes, token.Value)
		case TokenTypeAttribute:
			current.Attribute = parseAttribute(token.Value)
		case TokenTypeCombinator:
			if current.Tag != "" || current.ID != "" || len(current.Classes) > 0 {
				selectors = append(selectors, current)
				current = &Selector{Type: SelectorTypeCompound}
			}
		}
	}

	if current.Tag != "" || current.ID != "" || len(current.Classes) > 0 {
		selectors = append(selectors, current)
	}

	return selectors
}

func parseAttribute(attr string) *AttributeSelector {
	// Remove [ and ]
	attr = strings.Trim(attr, "[]")

	// Split by operators
	operators := []string{"~=", "|=", "^=", "$=", "*=", "="}
	for _, op := range operators {
		if parts := strings.Split(attr, op); len(parts) == 2 {
			return &AttributeSelector{
				Key:      strings.TrimSpace(parts[0]),
				Value:    strings.Trim(strings.TrimSpace(parts[1]), "\"'"),
				Operator: op,
			}
		}
	}

	return &AttributeSelector{
		Key: attr,
	}
}
