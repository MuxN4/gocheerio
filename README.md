<p align="center"> <img src="logo/logo.svg" alt="GoCheerio Logo" width="150"/> </p> <h1 align="center">GoCheerio</h1>

## Features (Planned)

- Full DOM Traversal and Manipulation
- CSS Selector Support
- Method Chaining
- Attribute Handling
- Streaming Selectors
- Channel-based API
- Memory Efficient
- Full CSS Selector Support

## Installation

```bash
go get github.com/MuxN4/gocheerio
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"
    "github.com/MuxN4/gocheerio"
)

func main() {
    // Parse HTML document
    doc, err := gocheerio.Load(`
        <html>
            <body>
                <div class="content">
                    <h1 id="title">Hello, World!</h1>
                    <p class="description">Welcome to GoCheerio</p>
                </div>
            </body>
        </html>
    `)
    if err != nil {
        log.Fatal(err)
    }

    // Find elements using CSS selectors
    titleText := doc.Find("#title").Text()
    fmt.Println("Title:", titleText)

    // Get attribute values
    content := doc.Find(".content")
    classAttr, exists := content.Attr("class")
    if exists {
        fmt.Println("Class:", classAttr)
    }

    // Iterate over elements
    doc.Find("p").Each(func(i int, sel Selection) {
        fmt.Printf("Paragraph %d: %s\n", i, sel.Text())
    })
}
```

## API Reference

### Document Methods

```go
Load(html string) (Document, error)
```
Creates a new Document from HTML content.

### Selection Methods

```go
Find(selector string) Selection
```
Finds all elements matching the given CSS selector.

```go
Html() (string, error)
```
Gets the HTML contents of the first element in the selection.

```go
Text() string
```
Gets the combined text contents of all matched elements.

```go
Attr(name string) (string, bool)
```
Gets the value of the specified attribute for the first element.

```go
Each(func(int, Selection))
```
Iterates over each element in the selection.

```go
Length() int
```
Returns the number of elements in the selection.

## CSS Selector Support

Currently supported selectors:
- Element selectors (`div`, `p`, `span`)
- ID selectors (`#myId`)
- Class selectors (`.myClass`)
- Attribute selectors:
  - `[attr]` - Has attribute
  - `[attr=value]` - Exact match
  - `[attr~=value]` - Contains word
  - `[attr|=value]` - Starts with value or value-
  - `[attr^=value]` - Starts with
  - `[attr$=value]` - Ends with
  - `[attr*=value]` - Contains substring

## Contributing

Contributions are welcome! Please feel free to submit pull requests or create issues for bugs and feature requests. For detailed guidelines on how to contribute, please refer to the [contribution.md](CONTRIBUTING.md) file.

## Acknowledgments

- Inspired by [Cheerio](https://github.com/cheeriojs/cheerio)
- Built on top of Go's excellent [x/net/html](https://pkg.go.dev/golang.org/x/net/html) package

## Documentation

Detailed documentation will be available as the project develops.
