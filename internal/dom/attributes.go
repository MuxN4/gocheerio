package dom

import "golang.org/x/net/html"

// GetAttribute returns the value of the specified attribute
func (n *Node) GetAttribute(name string) (string, bool) {
	for _, attr := range n.Node.Attr {
		if attr.Key == name {
			return attr.Val, true
		}
	}
	return "", false
}

// SetAttribute sets the value of the specified attribute
func (n *Node) SetAttribute(name, value string) {
	for i, attr := range n.Node.Attr {
		if attr.Key == name {
			n.Node.Attr[i].Val = value
			return
		}
	}
	n.Node.Attr = append(n.Node.Attr, html.Attribute{
		Key: name,
		Val: value,
	})
}

// RemoveAttribute removes the specified attribute
func (n *Node) RemoveAttribute(name string) {
	attrs := make([]html.Attribute, 0, len(n.Node.Attr))
	for _, attr := range n.Node.Attr {
		if attr.Key != name {
			attrs = append(attrs, attr)
		}
	}
	n.Node.Attr = attrs
}
