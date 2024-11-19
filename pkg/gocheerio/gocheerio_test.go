package gocheerio

import (
	"strings"
	"testing"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		name    string
		html    string
		wantErr bool
	}{
		{
			name:    "valid HTML",
			html:    "<html><body><div>Hello</div></body></html>",
			wantErr: false,
		},
		{
			name:    "empty HTML",
			html:    "",
			wantErr: false,
		},
		{
			name:    "malformed HTML",
			html:    "<div>unclosed",
			wantErr: false, // HTML parser is forgiving
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc, err := Load(tt.html)
			if (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if doc == nil {
				t.Error("Load() returned nil document")
			}
		})
	}
}

func TestDocument_Find(t *testing.T) {
	html := `
		<html>
			<body>
				<div class="container">
					<p id="first">First paragraph</p>
					<p class="highlight">Second paragraph</p>
				</div>
			</body>
		</html>
	`

	doc, err := Load(html)
	if err != nil {
		t.Fatalf("Failed to load HTML: %v", err)
	}

	tests := []struct {
		name     string
		selector string
		want     int // expected number of matches
	}{
		{
			name:     "find by tag",
			selector: "p",
			want:     2,
		},
		{
			name:     "find by id",
			selector: "#first",
			want:     1,
		},
		{
			name:     "find by class",
			selector: ".highlight",
			want:     1,
		},
		{
			name:     "find nested",
			selector: "div p",
			want:     2,
		},
		{
			name:     "empty selector",
			selector: "",
			want:     0,
		},
		{
			name:     "no matches",
			selector: ".nonexistent",
			want:     0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sel := doc.Find(tt.selector)
			if got := sel.Length(); got != tt.want {
				t.Errorf("Find() = %v matches, want %v", got, tt.want)
			}
		})
	}
}

func TestSelection_Methods(t *testing.T) {
	html := `
		<div class="container" data-test="value">
			<p>First <span>paragraph</span></p>
			<p>Second paragraph</p>
		</div>
	`

	doc, err := Load(html)
	if err != nil {
		t.Fatalf("Failed to load HTML: %v", err)
	}

	t.Run("Html", func(t *testing.T) {
		sel := doc.Find(".container")
		html, err := sel.Html()
		if err != nil {
			t.Errorf("Html() error = %v", err)
		}
		if !strings.Contains(html, "First") || !strings.Contains(html, "Second") {
			t.Errorf("Html() = %v, want content with 'First' and 'Second'", html)
		}
	})

	t.Run("Text", func(t *testing.T) {
		sel := doc.Find(".container")
		text := sel.Text()
		if !strings.Contains(text, "First") || !strings.Contains(text, "Second") {
			t.Errorf("Text() = %v, want text with 'First' and 'Second'", text)
		}
	})

	t.Run("Attr", func(t *testing.T) {
		sel := doc.Find(".container")
		if val, exists := sel.Attr("data-test"); !exists || val != "value" {
			t.Errorf("Attr() = (%v, %v), want (value, true)", val, exists)
		}
	})

	t.Run("Each", func(t *testing.T) {
		count := 0
		doc.Find("p").Each(func(i int, s Selection) {
			count++
		})
		if count != 2 {
			t.Errorf("Each() called %v times, want 2", count)
		}
	})

	t.Run("Find nested", func(t *testing.T) {
		sel := doc.Find(".container").Find("span")
		if sel.Length() != 1 {
			t.Errorf("Find() nested = %v matches, want 1", sel.Length())
		}
		text := sel.Text()
		if text != "paragraph" {
			t.Errorf("Text() of nested span = %v, want 'paragraph'", text)
		}
	})
}

func TestEdgeCases(t *testing.T) {
	tests := []struct {
		name string
		html string
		test func(*testing.T, Document)
	}{
		{
			name: "empty document",
			html: "",
			test: func(t *testing.T, doc Document) {
				if text := doc.Text(); text != "" {
					t.Errorf("Text() = %v, want empty string", text)
				}
			},
		},
		{
			name: "empty selection",
			html: "<div></div>",
			test: func(t *testing.T, doc Document) {
				sel := doc.Find(".nonexistent")
				if html, err := sel.Html(); html != "" || err != nil {
					t.Errorf("Html() = (%v, %v), want ('', nil)", html, err)
				}
				if text := sel.Text(); text != "" {
					t.Errorf("Text() = %v, want empty string", text)
				}
				if _, exists := sel.Attr("anything"); exists {
					t.Error("Attr() exists = true, want false")
				}
			},
		},
		{
			name: "nested text nodes",
			html: "<div>Hello <span>World</span>!</div>",
			test: func(t *testing.T, doc Document) {
				text := doc.Find("div").Text()
				if !strings.Contains(text, "Hello") || !strings.Contains(text, "World") {
					t.Errorf("Text() = %v, want 'Hello World!'", text)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc, err := Load(tt.html)
			if err != nil {
				t.Fatalf("Failed to load HTML: %v", err)
			}
			tt.test(t, doc)
		})
	}
}
