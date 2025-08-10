package termhyo

import (
	"fmt"
	"strconv"
)

// ANSI escape sequences for text formatting
const (
	// Text formatting
	AnsiReset     = "\x1b[0m"
	AnsiBold      = "\x1b[1m"
	AnsiDim       = "\x1b[2m"
	AnsiItalic    = "\x1b[3m"
	AnsiUnderline = "\x1b[4m"
	AnsiBlink     = "\x1b[5m"
	AnsiReverse   = "\x1b[7m"
	AnsiStrike    = "\x1b[9m"

	// Foreground colors (text colors)
	AnsiBlack   = "\x1b[30m"
	AnsiRed     = "\x1b[31m"
	AnsiGreen   = "\x1b[32m"
	AnsiYellow  = "\x1b[33m"
	AnsiBlue    = "\x1b[34m"
	AnsiMagenta = "\x1b[35m"
	AnsiCyan    = "\x1b[36m"
	AnsiWhite   = "\x1b[37m"

	// Bright foreground colors
	AnsiBrightBlack   = "\x1b[90m"
	AnsiBrightRed     = "\x1b[91m"
	AnsiBrightGreen   = "\x1b[92m"
	AnsiBrightYellow  = "\x1b[93m"
	AnsiBrightBlue    = "\x1b[94m"
	AnsiBrightMagenta = "\x1b[95m"
	AnsiBrightCyan    = "\x1b[96m"
	AnsiBrightWhite   = "\x1b[97m"

	// Background colors
	AnsiBgBlack   = "\x1b[40m"
	AnsiBgRed     = "\x1b[41m"
	AnsiBgGreen   = "\x1b[42m"
	AnsiBgYellow  = "\x1b[43m"
	AnsiBgBlue    = "\x1b[44m"
	AnsiBgMagenta = "\x1b[45m"
	AnsiBgCyan    = "\x1b[46m"
	AnsiBgWhite   = "\x1b[47m"

	// Bright background colors
	AnsiBgBrightBlack   = "\x1b[100m"
	AnsiBgBrightRed     = "\x1b[101m"
	AnsiBgBrightGreen   = "\x1b[102m"
	AnsiBgBrightYellow  = "\x1b[103m"
	AnsiBgBrightBlue    = "\x1b[104m"
	AnsiBgBrightMagenta = "\x1b[105m"
	AnsiBgBrightCyan    = "\x1b[106m"
	AnsiBgBrightWhite   = "\x1b[107m"
)

// HeaderStyle defines the styling for table headers
type HeaderStyle struct {
	// Text formatting
	Bold      bool
	Underline bool
	Italic    bool
	Dim       bool
	Blink     bool
	Reverse   bool
	Strike    bool

	// Colors
	ForegroundColor string // ANSI color code or empty for default
	BackgroundColor string // ANSI color code or empty for default

	// Custom escape sequences
	CustomPrefix string // Custom ANSI sequence to prepend
	CustomSuffix string // Custom ANSI sequence to append
}

// DefaultHeaderStyle returns a default header style with bold and underline
func DefaultHeaderStyle() HeaderStyle {
	return HeaderStyle{
		Bold:      true,
		Underline: true,
	}
}

// BoldHeaderStyle returns a header style with bold text
func BoldHeaderStyle() HeaderStyle {
	return HeaderStyle{
		Bold: true,
	}
}

// UnderlineHeaderStyle returns a header style with underlined text
func UnderlineHeaderStyle() HeaderStyle {
	return HeaderStyle{
		Underline: true,
	}
}

// ColoredHeaderStyle returns a header style with specified colors
func ColoredHeaderStyle(fgColor, bgColor string) HeaderStyle {
	return HeaderStyle{
		ForegroundColor: fgColor,
		BackgroundColor: bgColor,
	}
}

// ApplyStyle applies the header style to the given text
func (hs HeaderStyle) ApplyStyle(text string) string {
	if hs.isEmpty() {
		return text
	}

	var prefix string
	var suffix string = AnsiReset

	// Add custom prefix if specified
	if hs.CustomPrefix != "" {
		prefix += hs.CustomPrefix
	}

	// Add text formatting
	if hs.Bold {
		prefix += AnsiBold
	}
	if hs.Dim {
		prefix += AnsiDim
	}
	if hs.Italic {
		prefix += AnsiItalic
	}
	if hs.Underline {
		prefix += AnsiUnderline
	}
	if hs.Blink {
		prefix += AnsiBlink
	}
	if hs.Reverse {
		prefix += AnsiReverse
	}
	if hs.Strike {
		prefix += AnsiStrike
	}

	// Add foreground color
	if hs.ForegroundColor != "" {
		prefix += hs.ForegroundColor
	}

	// Add background color
	if hs.BackgroundColor != "" {
		prefix += hs.BackgroundColor
	}

	// Add custom suffix if specified
	if hs.CustomSuffix != "" {
		suffix = hs.CustomSuffix + suffix
	}

	return prefix + text + suffix
}

// getPrefix returns the ANSI prefix for the header style
func (hs HeaderStyle) getPrefix() string {
	if hs.isEmpty() {
		return ""
	}

	var prefix string

	// Add custom prefix if specified
	if hs.CustomPrefix != "" {
		prefix += hs.CustomPrefix
	}

	// Add text formatting
	if hs.Bold {
		prefix += AnsiBold
	}
	if hs.Dim {
		prefix += AnsiDim
	}
	if hs.Italic {
		prefix += AnsiItalic
	}
	if hs.Underline {
		prefix += AnsiUnderline
	}
	if hs.Blink {
		prefix += AnsiBlink
	}
	if hs.Reverse {
		prefix += AnsiReverse
	}
	if hs.Strike {
		prefix += AnsiStrike
	}

	// Add foreground color
	if hs.ForegroundColor != "" {
		prefix += hs.ForegroundColor
	}

	// Add background color
	if hs.BackgroundColor != "" {
		prefix += hs.BackgroundColor
	}

	return prefix
}

// getSuffix returns the ANSI suffix for the header style
func (hs HeaderStyle) getSuffix() string {
	if hs.isEmpty() {
		return ""
	}

	var suffix string = AnsiReset

	// Add custom suffix if specified
	if hs.CustomSuffix != "" {
		suffix = hs.CustomSuffix + suffix
	}

	return suffix
}

// isEmpty checks if the header style has any formatting applied
func (hs HeaderStyle) isEmpty() bool {
	return !hs.Bold && !hs.Underline && !hs.Italic && !hs.Dim &&
		!hs.Blink && !hs.Reverse && !hs.Strike &&
		hs.ForegroundColor == "" && hs.BackgroundColor == "" &&
		hs.CustomPrefix == "" && hs.CustomSuffix == ""
}

// Combine combines this header style with another, with the other style taking precedence
func (hs HeaderStyle) Combine(other HeaderStyle) HeaderStyle {
	result := hs

	// Other style takes precedence for boolean values if set
	if other.Bold {
		result.Bold = true
	}
	if other.Underline {
		result.Underline = true
	}
	if other.Italic {
		result.Italic = true
	}
	if other.Dim {
		result.Dim = true
	}
	if other.Blink {
		result.Blink = true
	}
	if other.Reverse {
		result.Reverse = true
	}
	if other.Strike {
		result.Strike = true
	}

	// Other style takes precedence for colors if set
	if other.ForegroundColor != "" {
		result.ForegroundColor = other.ForegroundColor
	}
	if other.BackgroundColor != "" {
		result.BackgroundColor = other.BackgroundColor
	}

	// Other style takes precedence for custom sequences if set
	if other.CustomPrefix != "" {
		result.CustomPrefix = other.CustomPrefix
	}
	if other.CustomSuffix != "" {
		result.CustomSuffix = other.CustomSuffix
	}

	return result
}

// RGB256 returns a 256-color ANSI code for foreground
func RGB256(colorCode int) string {
	return "\x1b[38;5;" + strconv.Itoa(colorCode) + "m"
}

// BgRGB256 returns a 256-color ANSI code for background
func BgRGB256(colorCode int) string {
	return "\x1b[48;5;" + strconv.Itoa(colorCode) + "m"
}

// TrueColorFg returns a true color (24-bit) ANSI code for foreground
func TrueColorFg(r, g, b int) string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b)
}

// TrueColorBg returns a true color (24-bit) ANSI code for background
func TrueColorBg(r, g, b int) string {
	return fmt.Sprintf("\x1b[48;2;%d;%d;%dm", r, g, b)
}
