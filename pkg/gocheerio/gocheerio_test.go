package gocheerio

import "testing"

func TestLoad(t *testing.T) {
	html := `<html><body><div id="test">Hello</div></body></html>`
	doc := Load(html)

	if doc == nil {
		t.Error("Load() returned nil")
	}

	rendered, err := doc.Html()
	if err != nil {
		t.Errorf("Html() returned error: %v", err)
	}
	if rendered == "" {
		t.Error("Html() returned empty string")
	}
}

func TestFind(t *testing.T) {
	html := `<html><body><div id="test">Hello</div></body></html>`
	doc := Load(html)

	sel := doc.Find("#test")
	if sel == nil {
		t.Error("Find() returned nil selection")
	}

	// Once selector matching is implemented, steve, add more specific tests
}
