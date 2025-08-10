/*
Package termhyo provides a flexible and feature-rich library for rendering beautiful tables in terminal applications.

The name combines "terminal" and the Japanese word "表 (hyo)" meaning "table".

# Features

termhyo offers comprehensive table rendering capabilities:

- Multiple rendering modes (Buffered and Streaming)
- Rich border styling options (Box Drawing, ASCII, Rounded, Markdown)
- Comprehensive header styling with ANSI escape sequences
- Unicode and multi-byte character support
- Automatic column width calculation
- Progressive border control
- Extensible renderer interface

# Basic Usage

Creating a simple table:

	columns := []termhyo.Column{
		{Title: "ID", Width: 0, Align: "right"},
		{Title: "Name", Width: 0, Align: "left"},
		{Title: "Age", Width: 0, Align: "center"},
	}

	table := termhyo.NewTable(os.Stdout, columns)
	table.AddRow("1", "Alice", "25")
	table.AddRow("2", "Bob", "30")
	table.Render()

# Header Styling

Adding visual styling to headers:

	headerStyle := termhyo.HeaderStyle{
		Bold:            true,
		ForegroundColor: termhyo.AnsiWhite,
		BackgroundColor: termhyo.AnsiBgBlue,
	}
	table.SetHeaderStyle(headerStyle)

# Border Control

Controlling table borders for different visual styles:

	// Remove header separator for clean look
	table.SetHeaderStyleWithoutSeparator(headerStyle)

	// Remove all horizontal borders
	table.SetHeaderStyleWithoutBorders(headerStyle)

	// Minimal table with only header styling
	table.SetHeaderStyleMinimal(headerStyle)

# Rendering Modes

termhyo automatically selects the appropriate rendering mode:

- Buffered Mode: Used when columns have auto-width (Width: 0). Calculates optimal widths.
- Streaming Mode: Used when all columns have fixed widths. More memory efficient.

# Character Support

Proper handling of various character types:

- ASCII characters
- Unicode characters
- Multi-byte characters (Japanese, Chinese, Korean)
- Combining characters
- Emoji and other special characters
- ANSI escape sequences

The library correctly calculates display width for all these character types.

# Color Support

Full ANSI color support for headers:

- 8-bit colors (AnsiRed, AnsiBlue, etc.)
- 256-color palette (RGB256)
- 24-bit true color (TrueColorFg, TrueColorBg)
- Text formatting (Bold, Italic, Underline, etc.)

# Examples

See the examples/ directory for comprehensive usage examples including:

- Basic table creation
- Header styling variations
- Border style demonstrations
- Unicode character handling
- Markdown table output
- Streaming mode usage
*/
package termhyo
