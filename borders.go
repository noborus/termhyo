package termhyo

// BorderStyle defines different border styles
type BorderStyle string

const (
	// BoxDrawingStyle uses Unicode box drawing characters
	BoxDrawingStyle BorderStyle = "box"
	// ASCIIStyle uses ASCII characters
	ASCIIStyle BorderStyle = "ascii"
	// RoundedStyle uses rounded corners
	RoundedStyle BorderStyle = "rounded"
	// DoubleStyle uses double line box drawing
	DoubleStyle BorderStyle = "double"
	// MinimalStyle uses minimal borders
	MinimalStyle BorderStyle = "minimal"
	// MarkdownStyle uses Markdown table format
	MarkdownStyle BorderStyle = "markdown"
)

// Predefined border character maps
var (
	boxDrawingBorders = map[string]string{
		"horizontal":   "─",
		"vertical":     "│",
		"cross":        "┼",
		"top_left":     "┌",
		"top_right":    "┐",
		"bottom_left":  "└",
		"bottom_right": "┘",
		"top_cross":    "┬",
		"bottom_cross": "┴",
		"left_cross":   "├",
		"right_cross":  "┤",
	}

	asciiBorders = map[string]string{
		"horizontal":   "-",
		"vertical":     "|",
		"cross":        "+",
		"top_left":     "+",
		"top_right":    "+",
		"bottom_left":  "+",
		"bottom_right": "+",
		"top_cross":    "+",
		"bottom_cross": "+",
		"left_cross":   "+",
		"right_cross":  "+",
	}

	roundedBorders = map[string]string{
		"horizontal":   "─",
		"vertical":     "│",
		"cross":        "┼",
		"top_left":     "╭",
		"top_right":    "╮",
		"bottom_left":  "╰",
		"bottom_right": "╯",
		"top_cross":    "┬",
		"bottom_cross": "┴",
		"left_cross":   "├",
		"right_cross":  "┤",
	}

	doubleBorders = map[string]string{
		"horizontal":   "═",
		"vertical":     "║",
		"cross":        "╬",
		"top_left":     "╔",
		"top_right":    "╗",
		"bottom_left":  "╚",
		"bottom_right": "╝",
		"top_cross":    "╦",
		"bottom_cross": "╩",
		"left_cross":   "╠",
		"right_cross":  "╣",
	}

	minimalBorders = map[string]string{
		"horizontal":   " ",
		"vertical":     " ",
		"cross":        " ",
		"top_left":     " ",
		"top_right":    " ",
		"bottom_left":  " ",
		"bottom_right": " ",
		"top_cross":    " ",
		"bottom_cross": " ",
		"left_cross":   " ",
		"right_cross":  " ",
	}

	markdownBorders = map[string]string{
		"horizontal":   "-",
		"vertical":     "|",
		"cross":        "|",
		"top_left":     "",
		"top_right":    "",
		"bottom_left":  "",
		"bottom_right": "",
		"top_cross":    "|",
		"bottom_cross": "|",
		"left_cross":   "|",
		"right_cross":  "|",
	}
)

// getBorderChars returns border characters for the specified style
func getBorderChars(style BorderStyle) map[string]string {
	switch style {
	case ASCIIStyle:
		return asciiBorders
	case RoundedStyle:
		return roundedBorders
	case DoubleStyle:
		return doubleBorders
	case MinimalStyle:
		return minimalBorders
	case MarkdownStyle:
		return markdownBorders
	default: // BoxDrawingStyle
		return boxDrawingBorders
	}
}
