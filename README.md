# termhyo

[![Go Reference](https://pkg.go.dev/badge/github.com/noborus/termhyo.svg)](https://pkg.go.dev/github.com/noborus/termhyo)
[![Go Report Card](https://goreportcard.com/badge/github.com/noborus/termhyo)](https://goreportcard.com/report/github.com/noborus/termhyo)
[![GitHub release](https://img.shields.io/github/release/noborus/termhyo.svg)](https://github.com/noborus/termhyo/releases)
[![License](https://img.shields.io/github/license/noborus/termhyo.svg)](LICENSE)

`termhyo` is a Go package for beautifully displaying tabular data. The name combines "terminal" and the Japanese word "表 (hyo)" meaning "table", and is specialized for terminal display.

## Features

- **Two rendering modes**: Flexible display with BufferedMode and StreamingMode
- **Multiple border styles**: Choose from Box Drawing, ASCII, Rounded, Double, Minimal, and VerticalBar (only vertical | separators)
- **Automatic width calculation**: Automatic column width adjustment and alignment
- **Unicode support**: Proper handling of multibyte characters, combining characters, emojis, and East Asian text
- **Interface design**: Extensible renderer architecture

## Installation

To install termhyo, use `go get`:

```bash
go get github.com/noborus/termhyo
```

### Requirements

- Go 1.23 or later

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
        {Title: "ID", Width: 0, Align: termhyo.Right},
        {Title: "Name", Width: 0, Align: termhyo.Left},
        {Title: "Score", Width: 0, Align: termhyo.Center},
    }

    table := termhyo.NewTable(os.Stdout, columns)
    table.AddRow("1", "Alice", "85")
    table.AddRow("2", "Bob", "92")
    table.Render()
}
```

### Changing Border Style

```go
table := termhyo.NewTable(os.Stdout, columns, termhyo.Border(termhyo.ASCIIStyle))
```

### Text Alignment

termhyo provides type-safe alignment options:

```go
// Available alignment constants
termhyo.Left     // Left-aligned text
termhyo.Center   // Center-aligned text
termhyo.Right    // Right-aligned text
termhyo.Default  // Default/unspecified alignment (defaults to left)

// Column-level alignment
columns := []termhyo.Column{
    {Title: "ID", Align: termhyo.Right},
    {Title: "Name", Align: termhyo.Left},
    {Title: "Score", Align: termhyo.Center},
}

// Cell-level alignment (overrides column alignment)
table.AddRowCells(
    termhyo.Cell{Content: "1", Align: termhyo.Center},
    termhyo.Cell{Content: "Alice"},  // Uses column alignment
    termhyo.Cell{Content: "85"},
)
```

### Custom Border Configuration

```go
// Create custom border configuration
customConfig := termhyo.TableBorderConfig{
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

table := termhyo.NewTable(os.Stdout, columns, termhyo.BorderConfig(customConfig))
```

## Running Examples

You can run various example programs to see termhyo in action.
See the [examples](./examples) directory for more.

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
- `VerticalBarStyle`: Only vertical bar separators (|), no outer borders
- `MarkdownStyle`: Markdown table format
- `TSVStyle`: Tab-separated values format

## License

MIT License
