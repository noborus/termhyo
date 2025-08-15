/*
Package termhyo provides a flexible and feature-rich library for rendering beautiful tables in terminal applications.

The name combines "terminal" and the Japanese word "表 (hyo)" meaning "table".

# Features

- Multiple rendering modes: Buffered and Streaming
- Rich border styling: BoxDrawing, ASCII, Rounded, Double, Minimal, VerticalBar (only |), Markdown, TSV, and custom borders
- Functional Option Pattern for table configuration (Border, Header, AutoAlign, BorderConfig, etc.)
- Comprehensive header styling with ANSI escape sequences and true color
- Unicode and multi-byte character support (East Asian, combining, emoji, etc.)
- Automatic column width calculation and type-safe alignment
- Fine-grained border control (show/hide top, bottom, left, right, vertical, middle)
- Extensible renderer interface

# Basic Usage

	columns := []termhyo.Column{
		{Title: "ID", Align: termhyo.Right},
		{Title: "Name", Align: termhyo.Left},
		{Title: "Score", Align: termhyo.Center},
	}
	table := termhyo.NewTable(os.Stdout, columns)
	table.AddRow("1", "Alice", "85")
	table.AddRow("2", "Bob", "92")
	table.Render()

# Border Style Example

	table := termhyo.NewTable(os.Stdout, columns, termhyo.Border(termhyo.VerticalBarStyle), termhyo.AutoAlign(false))

# Header Styling Example

	headerStyle := termhyo.HeaderStyle{
		Bold:            true,
		ForegroundColor: termhyo.AnsiWhite,
		BackgroundColor: termhyo.AnsiBgBlue,
	}
	table.SetHeaderStyle(headerStyle)

# Custom Border Example

	customConfig := termhyo.TableBorderConfig{
		Chars: map[string]string{"vertical": "|"},
		Top: false, Bottom: false, Left: false, Right: false, Vertical: true,
	}
	table := termhyo.NewTable(os.Stdout, columns, termhyo.BorderConfig(customConfig))

# Rendering Modes

- BufferedMode: Used when columns have auto-width (Width: 0) or alignment is enabled. Calculates optimal widths.
- StreamingMode: Used when all columns have fixed widths and alignment is disabled. Memory efficient.

# Character & Color Support

- Handles ASCII, Unicode, multi-byte, combining, emoji, and ANSI escape sequences
- Correct display width calculation for all character types
- Full ANSI color, 256-color, and 24-bit true color for headers
- Text formatting: Bold, Italic, Underline, Reverse, etc.

# See Also

See the examples/ directory for comprehensive usage:

- Basic table creation
- Header styling variations
- Border style demonstrations (including VerticalBarStyle)
- Unicode and multilingual support
- Markdown/TSV output
- Streaming mode usage
- Custom border configuration
*/
package termhyo
