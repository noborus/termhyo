package termhyo

import (
	"regexp"
	"strings"

	"github.com/rivo/uniseg"
)

// ANSI escape sequence patterns.
var (
	// ANSI color codes and other escape sequences.
	ansiEscapeRegex = regexp.MustCompile(`\x1b\[[0-9;]*[a-zA-Z]`)
	// Other control sequences (like \r, \n, \t etc.).
	controlCharsRegex = regexp.MustCompile(`[\x00-\x1f\x7f]`)
)

// stripEscapeSequences removes ANSI escape sequences and control characters.
func stripEscapeSequences(s string) string {
	// Remove ANSI escape sequences
	s = ansiEscapeRegex.ReplaceAllString(s, "")
	// Remove other control characters except for normal whitespace
	s = controlCharsRegex.ReplaceAllStringFunc(s, func(match string) string {
		// Keep normal spaces and tabs, remove others
		if match == " " || match == "\t" {
			return match
		}
		return ""
	})
	return s
}

// StringWidth returns the display width of a string on terminal.
// This properly handles multibyte characters, combining characters, emojis, and ANSI escape sequences.
// This is the public version of stringWidth for external use.
func StringWidth(s string) int {
	return stringWidth(s)
}

// stringWidth returns the display width of a string on terminal.
// This properly handles multibyte characters, combining characters, emojis, and ANSI escape sequences.
func stringWidth(s string) int {
	// First remove escape sequences and control characters
	cleaned := stripEscapeSequences(s)
	return uniseg.StringWidth(cleaned)
}

// truncateString truncates a string to fit within the specified display width.
// Preserves ANSI escape sequences while calculating display width correctly.
func truncateString(s string, maxWidth int) string {
	if maxWidth <= 0 {
		return ""
	}

	// Check if truncation is needed by comparing display width
	if stringWidth(s) <= maxWidth {
		return s
	}

	// If we need ellipsis, reserve space for it
	ellipsisWidth := 3
	if maxWidth < ellipsisWidth {
		ellipsisWidth = maxWidth
		if ellipsisWidth == 0 {
			return ""
		}
	}

	targetWidth := maxWidth - ellipsisWidth
	if targetWidth <= 0 {
		return strings.Repeat(".", maxWidth)
	}

	// Split the string to handle escape sequences properly
	result := truncateWithEscapes(s, targetWidth)
	return result + "..."
}

// truncateWithEscapes truncates a string while preserving escape sequences.
func truncateWithEscapes(s string, maxWidth int) string {
	if maxWidth <= 0 {
		return ""
	}

	var result strings.Builder
	var currentWidth int
	gr := uniseg.NewGraphemes(s)

	for gr.Next() {
		cluster := gr.Str()
		// Detect ANSI escape sequence
		if strings.HasPrefix(cluster, "\x1b[") {
			result.WriteString(cluster)
			continue
		}
		// Detect control characters
		runes := []rune(cluster)
		if len(runes) == 1 && (runes[0] < 0x20 || runes[0] == 0x7f) {
			if runes[0] == ' ' || runes[0] == '\t' {
				if currentWidth >= maxWidth {
					break
				}
				result.WriteString(cluster)
				currentWidth++
			}
			continue
		}
		// Normal character
		clusterWidth := uniseg.StringWidth(cluster)
		if currentWidth+clusterWidth > maxWidth {
			break
		}
		result.WriteString(cluster)
		currentWidth += clusterWidth
	}

	return result.String()
}

// padString pads a string to the specified display width with spaces.
// Correctly handles ANSI escape sequences when calculating padding.
func padString(s string, width int, align Alignment) string {
	currentWidth := stringWidth(s) // This now handles escape sequences correctly
	if currentWidth >= width {
		return s
	}

	padding := width - currentWidth

	switch align {
	case Right:
		return spaces(padding) + s
	case Center:
		leftPad := padding / 2
		rightPad := padding - leftPad
		return spaces(leftPad) + s + spaces(rightPad)
	default: // Left or Default
		return s + spaces(padding)
	}
}

// spaces returns a string with n spaces.
func spaces(n int) string {
	if n <= 0 {
		return ""
	}
	result := make([]byte, n)
	for i := range result {
		result[i] = ' '
	}
	return string(result)
}
