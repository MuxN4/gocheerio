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
    "github.com/MuxN4/gocheerio"
)

func main() {
    // Example code will be added as features are implemented
    doc := gocheerio.Load("<html><body><div class=\"example\">Hello, World!</div></body></html>")
    text := doc.Find(".example").Text()
    fmt.Println(text)
}
```

## Documentation

Detailed documentation will be available as the project develops.
