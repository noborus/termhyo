package termhyo

// BorderStyle defines different border styles
type BorderStyle string

// BorderConfig holds border style configuration
type BorderConfig struct {
	Chars           map[string]string
	DisableTop      bool
	DisableBottom   bool
	DisableMiddle   bool
	DisableLeft     bool
	DisableRight    bool
	DisableVertical bool // Controls internal vertical separators
	DisablePadding  bool // Controls content padding
}

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
	// TSVStyle uses tab separators only
	TSVStyle BorderStyle = "tsv"
)

// Predefined border configurations
var (
	boxDrawingConfig = BorderConfig{
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
		DisableTop:      false,
		DisableBottom:   false,
		DisableMiddle:   false,
		DisableLeft:     false,
		DisableRight:    false,
		DisableVertical: false,
		DisablePadding:  false,
	}

	asciiConfig = BorderConfig{
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
		DisableTop:      false,
		DisableBottom:   false,
		DisableMiddle:   false,
		DisableLeft:     false,
		DisableRight:    false,
		DisableVertical: false,
		DisablePadding:  false,
	}

	roundedConfig = BorderConfig{
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
		DisableTop:      false,
		DisableBottom:   false,
		DisableMiddle:   false,
		DisableLeft:     false,
		DisableRight:    false,
		DisableVertical: false,
		DisablePadding:  false,
	}

	doubleConfig = BorderConfig{
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
		DisableTop:      false,
		DisableBottom:   false,
		DisableMiddle:   false,
		DisableLeft:     false,
		DisableRight:    false,
		DisableVertical: false,
		DisablePadding:  false,
	}

	minimalConfig = BorderConfig{
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
		DisableTop:      true,
		DisableBottom:   true,
		DisableMiddle:   true,
		DisableLeft:     true,
		DisableRight:    true,
		DisableVertical: true,
		DisablePadding:  false,
	}

	markdownConfig = BorderConfig{
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
		DisableTop:      true,
		DisableBottom:   true,
		DisableMiddle:   false,
		DisableLeft:     false,
		DisableRight:    false,
		DisableVertical: false,
		DisablePadding:  false,
	}

	tsvConfig = BorderConfig{
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
		DisableTop:      true,
		DisableBottom:   true,
		DisableMiddle:   true,
		DisableLeft:     true,
		DisableRight:    true,
		DisableVertical: false,
		DisablePadding:  true, // TSVではパディングを無効化
	}
)

// getBorderConfig returns border configuration for the specified style
func getBorderConfig(style BorderStyle) BorderConfig {
	switch style {
	case ASCIIStyle:
		return asciiConfig
	case RoundedStyle:
		return roundedConfig
	case DoubleStyle:
		return doubleConfig
	case MinimalStyle:
		return minimalConfig
	case MarkdownStyle:
		return markdownConfig
	case TSVStyle:
		return tsvConfig
	default: // BoxDrawingStyle
		return boxDrawingConfig
	}
}

// getBorderChars returns border characters for the specified style (for backward compatibility)
func getBorderChars(style BorderStyle) map[string]string {
	return getBorderConfig(style).Chars
}
