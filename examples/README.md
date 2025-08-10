# termhyo Examples

This directory contains various examples demonstrating the features of termhyo.

## Running Examples

Each example is in its own directory with a `main.go` file. You can run them in two ways:

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

## Available Examples

### [basic/](basic/)
Basic table creation and rendering. Good starting point.

**Features:** Column definition, data addition, basic rendering

### [header_styles/](header_styles/)
Advanced header styling with ANSI escape sequences.

**Features:** Colors, text formatting, border control, convenience methods

### [header_full_line/](header_full_line/)
Full-line header styling examples.

**Features:** Complete line styling, background colors, visual distinction

### [streaming/](streaming/)
Streaming mode for large datasets with fixed-width columns.

**Features:** Memory-efficient rendering, fixed-width columns

### [japanese/](japanese/)
Japanese text handling and proper width calculation.

**Features:** Multi-byte characters, proper alignment, Unicode support

### [unicode/](unicode/)
Unicode character support including emojis and special characters.

**Features:** Emoji, combining characters, various Unicode blocks

### [combining/](combining/)
Combining character handling and width calculation.

**Features:** Diacritics, combining marks, proper display width

### [markdown/](markdown/)
Markdown table format output.

**Features:** Markdown syntax, alignment indicators, header styling

### [custom_borders/](custom_borders/)
Custom border configuration and styling.

**Features:** Border customization, selective border removal

### [styles/](styles/)
Different border styles demonstration.

**Features:** Box drawing, ASCII, rounded borders

## Learning Path

1. **Start with [basic/](basic/)** - Learn fundamental concepts
2. **Try [header_styles/](header_styles/)** - Understand styling options
3. **Explore [unicode/](unicode/) and [japanese/](japanese/)** - Character handling
4. **Check [streaming/](streaming/)** - Performance considerations
5. **Review [markdown/](markdown/)** - Alternative output formats
6. **Experiment with [custom_borders/](custom_borders/)** - Advanced customization

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
