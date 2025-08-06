# termhyo

`termhyo` is a Go package for beautifully displaying tabular data. The name combines "terminal" and the Japanese word "表 (hyo)" meaning "table", and is specialized for terminal display.

## Features

- **Two rendering modes**: Flexible display with BufferedMode and StreamingMode
- **Multiple border styles**: Choose from Box Drawing, ASCII, Rounded, Double, and Minimal
- **Automatic width calculation**: Automatic column width adjustment and alignment
- **Interface design**: Extensible renderer architecture

## File Structure

```tree
termhyo/
├── writer.go          # Package documentation and main entry point
├── table.go           # Table struct and main logic
├── column.go          # Column, Cell, Row definitions
├── borders.go         # Border style definitions
├── renderer.go        # Renderer interface and implementation
├── examples.go        # Usage examples
└── examples/          # Runnable example programs
    ├── basic.go       # Basic table example
    ├── styles.go      # Border styles demonstration
    └── streaming.go   # Streaming mode example
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

## License

MIT License
