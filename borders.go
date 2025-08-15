package termhyo

// BorderStyle defines different border styles.
type BorderStyle string

// TableBorderConfig holds border style configuration.
type TableBorderConfig struct {
	Chars    map[string]string
	Top      bool // Show top border
	Bottom   bool // Show bottom border
	Middle   bool // Show middle separator between header and data
	Left     bool // Show left border
	Right    bool // Show right border
	Vertical bool // Show internal vertical separators
	Padding  bool // Add content padding
}

const (
	// BoxDrawingStyle uses Unicode box drawing characters.
	BoxDrawingStyle BorderStyle = "box"
	// ASCIIStyle uses ASCII characters.
	ASCIIStyle BorderStyle = "ascii"
	// RoundedStyle uses rounded corners.
	RoundedStyle BorderStyle = "rounded"
	// DoubleStyle uses double line box drawing.
	DoubleStyle BorderStyle = "double"
	// MinimalStyle uses minimal borders.
	MinimalStyle BorderStyle = "minimal"
	// VerticalBarStyle uses only vertical bar separators (|), no outer borders.
	VerticalBarStyle BorderStyle = "vertical_bar"
	// MarkdownStyle uses Markdown table format.
	MarkdownStyle BorderStyle = "markdown"
	// TSVStyle uses tab separators only.
	TSVStyle BorderStyle = "tsv"
)

// Predefined border configurations.
var (
	boxDrawingConfig = TableBorderConfig{
		Chars: map[string]string{
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
		},
		Top:      true,
		Bottom:   true,
		Middle:   true,
		Left:     true,
		Right:    true,
		Vertical: true,
		Padding:  true,
	}

	asciiConfig = TableBorderConfig{
		Chars: map[string]string{
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
		},
		Top:      true,
		Bottom:   true,
		Middle:   true,
		Left:     true,
		Right:    true,
		Vertical: true,
		Padding:  true,
	}

	roundedConfig = TableBorderConfig{
		Chars: map[string]string{
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
		},
		Top:      true,
		Bottom:   true,
		Middle:   true,
		Left:     true,
		Right:    true,
		Vertical: true,
		Padding:  true,
	}

	doubleConfig = TableBorderConfig{
		Chars: map[string]string{
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
		},
		Top:      true,
		Bottom:   true,
		Middle:   true,
		Left:     true,
		Right:    true,
		Vertical: true,
		Padding:  true,
	}

	verticalBarConfig = TableBorderConfig{
		Chars: map[string]string{
			"horizontal":   "",
			"vertical":     "|",
			"cross":        "|",
			"top_left":     "",
			"top_right":    "",
			"bottom_left":  "",
			"bottom_right": "",
			"top_cross":    "",
			"bottom_cross": "",
			"left_cross":   "",
			"right_cross":  "",
		},
		Top:      false,
		Bottom:   false,
		Middle:   false,
		Left:     false,
		Right:    false,
		Vertical: true,
		Padding:  true,
	}

	minimalConfig = TableBorderConfig{
		Chars: map[string]string{
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
		},
		Top:      false,
		Bottom:   false,
		Middle:   false,
		Left:     false,
		Right:    false,
		Vertical: false,
		Padding:  true,
	}

	markdownConfig = TableBorderConfig{
		Chars: map[string]string{
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
		},
		Top:      false,
		Bottom:   false,
		Middle:   true,
		Left:     true,
		Right:    true,
		Vertical: true,
		Padding:  true, // Use common padding functionality
	}

	tsvConfig = TableBorderConfig{
		Chars: map[string]string{
			"horizontal":   "",
			"vertical":     "\t",
			"cross":        "",
			"top_left":     "",
			"top_right":    "",
			"bottom_left":  "",
			"bottom_right": "",
			"top_cross":    "",
			"bottom_cross": "",
			"left_cross":   "",
			"right_cross":  "",
		},
		Top:      false,
		Bottom:   false,
		Middle:   false,
		Left:     false,
		Right:    false,
		Vertical: true,
		Padding:  false, // Disable padding for TSV format
	}
)

// GetBorderConfig returns border configuration for the specified style.
func GetBorderConfig(style BorderStyle) TableBorderConfig {
	switch style {
	case ASCIIStyle:
		return asciiConfig
	case RoundedStyle:
		return roundedConfig
	case DoubleStyle:
		return doubleConfig
	case MinimalStyle:
		return minimalConfig
	case VerticalBarStyle:
		return verticalBarConfig
	case MarkdownStyle:
		return markdownConfig
	case TSVStyle:
		return tsvConfig
	default: // BoxDrawingStyle
		return boxDrawingConfig
	}
}
