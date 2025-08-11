# termhyo

[![Go Reference](https://pkg.go.dev/badge/github.com/noborus/termhyo.svg)](https://pkg.go.dev/github.com/noborus/termhyo)
[![Go Report Card](https://goreportcard.com/badge/github.com/noborus/termhyo)](https://goreportcard.com/report/github.com/noborus/termhyo)
[![GitHub release](https://img.shields.io/github/release/noborus/termhyo.svg)](https://github.com/noborus/termhyo/releases)
[![License](https://img.shields.io/github/license/noborus/termhyo.svg)](LICENSE)

`termhyo` is a Go package for beautifully displaying tabular data. The name combines "terminal" and the Japanese word "表 (hyo)" meaning "table", and is specialized for terminal display.

## Features

- **Two rendering modes**: Flexible display with BufferedMode and StreamingMode
- **Multiple border styles**: Choose from Box Drawing, ASCII, Rounded, Double, and Minimal
- **Automatic width calculation**: Automatic column width adjustment and alignment
- **Unicode support**: Proper handling of multibyte characters, combining characters, emojis, and East Asian text
- **Interface design**: Extensible renderer architecture

## Installation

To install termhyo, use `go get`:

```bash
go get github.com/noborus/termhyo
```

### Requirements

- Go 1.21 or later

## File Structure

```tree
termhyo/
├── termhyo.go         # Package documentation and main entry point
├── table.go           # Table struct and main logic
├── column.go          # Column, Cell, Row definitions
├── borders.go         # Border style definitions
├── renderer.go        # Renderer interface and implementation
├── markdown.go        # Markdown table renderer
├── header_styles.go   # Header styling with ANSI escape sequences
├── width.go           # String width calculation utilities
└── examples/          # Runnable example programs
    ├── basic/         # Basic table example
    │   └── main.go
    ├── header_styles/ # Header styling demonstration
    │   └── main.go
    ├── streaming/     # Streaming mode example
    │   └── main.go
    ├── japanese/      # Japanese text example
    │   └── main.go
    ├── unicode/       # Unicode and emoji example
    │   └── main.go
    ├── combining/     # Combining characters example
    │   └── main.go
    ├── markdown/      # Markdown table format example
    │   └── main.go
    └── custom_borders/ # Custom border configuration example
        └── main.go
```

## Basic Usage

### Simple Table

```go
package main

import (
    "os"
    "github.com/noborus/termhyo"
)

func main() {
    columns := []termhyo.Column{
        {Title: "ID", Width: 0, Align: "right"},
        {Title: "Name", Width: 0, Align: "left"},
        {Title: "Score", Width: 0, Align: "center"},
    }

    table := termhyo.NewTable(os.Stdout, columns)
    table.AddRow("1", "Alice", "85")
    table.AddRow("2", "Bob", "92")
    table.Render()
}
```

### Changing Border Style

```go
table := termhyo.NewTableWithStyle(os.Stdout, columns, termhyo.ASCIIStyle)
```

### Custom Border Configuration

```go
// Create custom border configuration
customConfig := termhyo.BorderConfig{
    Chars: map[string]string{
        "horizontal": "=",
        "vertical":   "|",
        "cross":      "+",
        // ... other border characters
    },
    Top:      false, // No top border
    Bottom:   false, // No bottom border
    Middle:   true,  // Keep header separator
    Left:     false, // No left border
    Right:    false, // No right border
    Vertical: true,  // Keep internal column separators
    Padding:  true,  // Enable content padding
}

table.SetBorderConfig(customConfig)
```

## Running Examples

You can run the example programs to see termhyo in action:

```bash
# Basic table example
cd examples
go run basic.go

# Different border styles
go run styles.go

# Streaming mode demonstration
go run streaming.go

# Japanese text handling
go run japanese.go

# Unicode and emoji support
go run unicode.go

# Combining characters and complex Unicode
go run combining.go

# Markdown table format
go run markdown.go

# Custom border configurations
go run custom_borders.go
```

## Rendering Modes

### BufferedMode

- Collects all rows and renders them in batch
- Automatic width calculation and alignment possible
- Automatically selected when column width is 0 (auto) or alignment is enabled

### StreamingMode

- Renders immediately as rows are added
- Automatically selected when fixed width and alignment disabled
- Memory efficient

## Border Styles

- `BoxDrawingStyle`: Unicode Box Drawing characters (default)
- `ASCIIStyle`: ASCII characters
- `RoundedStyle`: Rounded corner style
- `DoubleStyle`: Double line style
- `MinimalStyle`: Minimal border
- `MarkdownStyle`: Markdown table format
- `TSVStyle`: Tab-separated values format

## License

MIT License
