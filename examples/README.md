# termhyo Examples

This directory contains various examples demonstrating the features and flexibility of termhyo.

## Running Examples

Each example is in its own directory with a `main.go` file. You can run them as follows:

### Method 1: Direct execution

```bash
cd basic
go run main.go
```

### Method 2: Build then run

```bash
cd basic
go build -o basic .
./basic
```

### Method 3: Using Makefile (from project root)

```bash
make examples  # Build all examples
```

## Example List

### [basic/](basic/)

Basic table creation and rendering. The best starting point for learning termhyo.

### [autoalign_false/](autoalign_false/)

Demonstrates `AutoAlign(false)` and `VerticalBarStyle`. Disables global auto-alignment and uses only vertical separators for a minimal look.

### [header_styles/](header_styles/)

Advanced header styling with ANSI colors, text decorations, border control, and convenience methods.

### [header_full_line/](header_full_line/)

Applies color and decoration to the entire header line.

### [streaming/](streaming/)

Streaming rendering for large datasets with fixed-width columns. Memory efficient output.

### [japanese/](japanese/)

Proper width calculation and alignment for Japanese (multibyte) text.

### [unicode/](unicode/)

Display of various Unicode characters, including emoji and special symbols.

### [combining/](combining/)

Handling and width calculation for combining characters (diacritics, etc).

### [markdown/](markdown/)

Table output in Markdown format.

### [custom_borders/](custom_borders/)

Custom border configuration. Flexible control, such as only internal separators or no borders at all.

### [styles/](styles/)

Comparison of border styles (BoxDrawing, ASCII, Rounded, Double, Minimal, VerticalBar, etc).

## Learning Path

1. **Start with [basic/](basic/)** - Learn the fundamentals
1. **Try [styles/](styles/) and [custom_borders/](custom_borders/)** - Explore border styles and customization
1. **Explore [header_styles/](header_styles/) and [header_full_line/](header_full_line/)** - Learn about header decoration and coloring
1. **Check [autoalign_false/](autoalign_false/)** - See minimal tables and alignment control
1. **Review [japanese/](japanese/), [unicode/](unicode/), and [combining/](combining/)** - Understand multilingual and Unicode support
1. **Experiment with [markdown/](markdown/) and [streaming/](streaming/)** - Use Markdown output and streaming rendering
1. **Try [header_styles/](header_styles/)** - Understand styling options
1. **Explore [unicode/](unicode/) and [japanese/](japanese/)** - Character handling
1. **Check [streaming/](streaming/)** - Performance considerations
1. **Review [markdown/](markdown/)** - Alternative output formats
1. **Experiment with [custom_borders/](custom_borders/)** - Advanced customization

## Building All Examples

From the project root:

```bash
make examples
```

This will build all examples into their respective directories.

## Clean Up

To remove built binaries:

```bash
make clean
```
