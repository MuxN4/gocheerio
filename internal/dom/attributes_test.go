package dom

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNodeGetAttribute(t *testing.T) {
	testCases := []struct {
		name    string
		attr    string
		want    string
		isExist bool
	}{
		{
			name:    "Try to Get A Tag Href Succesfully",
			attr:    "href",
			want:    "https://example.com",
			isExist: true,
		},
		{
			name:    "Try to Get Not Exist Attribute",
			attr:    "tag",
			want:    "",
			isExist: false,
		},
	}

	htmlContent := `<a href="https://example.com">example url</a>`

	doc, err := NewDocument(htmlContent)
	if err != nil {
		t.Fatalf("Failed to create document: %v", err)
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			body := doc.Root().FirstChild().NextSibling()
			a := body.FirstChild()
			attr, isExist := a.GetAttribute(tc.attr)
			if tc.isExist != isExist {
				t.Fatalf("Failed to find attribute %s", tc.attr)
			}
			assert.Equal(t, tc.want, attr)
		})
	}
}

func TestNodeSetAttribute(t *testing.T) {
	testCases := []struct {
		name string
		attr string
		want string
	}{
		{
			name: "Try to Set Existing Attribute",
			attr: "width",
			want: "100",
		},
		{
			name: "Try to Set Not Exist Attribute",
			attr: "height",
			want: "100",
		},
	}
	htmlContent := `<img src="https://example.com" width="100"/>`

	doc, err := NewDocument(htmlContent)
	if err != nil {
		t.Fatalf("Failed to create document: %v", err)
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			body := doc.Root().FirstChild().NextSibling()
			img := body.FirstChild()
			img.SetAttribute(tc.attr, tc.want)
			attr, isExist := img.GetAttribute(tc.attr)
			if !isExist {
				t.Fatalf("Failed to find attribute %s", tc.attr)
			}
			assert.Equal(t, tc.want, attr)
		})
	}
}

func TestNodeRemoveAttribute(t *testing.T) {
	testCases := []struct {
		name string
		attr string
		want string
	}{
		{
			name: "Try to Remove Existing Attribute",
			attr: "width",
		},
		{
			name: "Try to Remove Not Exist Attribute",
			attr: "height",
		},
	}
	htmlContent := `<img src="https://example.com" width="100"/>`

	doc, err := NewDocument(htmlContent)
	if err != nil {
		t.Fatalf("Failed to create document: %v", err)
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			body := doc.Root().FirstChild().NextSibling()
			img := body.FirstChild()
			img.RemoveAttribute(tc.attr)
			_, isExist := img.GetAttribute(tc.attr)
			if isExist {
				t.Fatalf("Failed to remove attribute %s", tc.attr)
			}
		})
	}
}
