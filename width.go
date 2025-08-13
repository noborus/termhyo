package termhyo

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/mattn/go-runewidth"
	"golang.org/x/text/unicode/norm"
)

// ANSI escape sequence patterns.
var (
	// ANSI color codes and other escape sequences
	ansiEscapeRegex = regexp.MustCompile(`\x1b\[[0-9;]*[a-zA-Z]`)
	// Other control sequences (like \r, \n, \t etc.)
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

	// Normalize the string to handle combining characters properly
	normalized := norm.NFC.String(cleaned)
	return runewidth.StringWidth(normalized)
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

	// Process the string character by character, handling escape sequences
	i := 0
	runes := []rune(s)

	for i < len(runes) {
		r := runes[i]

		// Check for ANSI escape sequence
		if r == '\x1b' && i+1 < len(runes) && runes[i+1] == '[' {
			// Find the end of the escape sequence
			escapeStart := i
			i += 2 // Skip \x1b[

			for i < len(runes) && !((runes[i] >= 'a' && runes[i] <= 'z') || (runes[i] >= 'A' && runes[i] <= 'Z')) {
				i++
			}
			if i < len(runes) {
				i++ // Include the final letter
			}

			// Add the entire escape sequence to result (doesn't count toward width)
			result.WriteString(string(runes[escapeStart:i]))
			continue
		}

		// Handle control characters
		if r < 0x20 || r == 0x7f {
			if r == ' ' || r == '\t' {
				// Count spaces and tabs toward width
				if currentWidth >= maxWidth {
					break
				}
				result.WriteRune(r)
				currentWidth++
			}
			// Skip other control characters
			i++
			continue
		}

		// Handle combining characters
		if unicode.Is(unicode.Mn, r) || unicode.Is(unicode.Me, r) || unicode.Is(unicode.Mc, r) {
			result.WriteRune(r)
			i++
			continue
		}

		// Regular character - check if it fits
		charWidth := runewidth.RuneWidth(r)
		if currentWidth+charWidth > maxWidth {
			break
		}

		result.WriteRune(r)
		currentWidth += charWidth
		i++
	}

	return result.String()
}

// padString pads a string to the specified display width with spaces.
// Correctly handles ANSI escape sequences when calculating padding.
func padString(s string, width int, align string) string {
	currentWidth := stringWidth(s) // This now handles escape sequences correctly
	if currentWidth >= width {
		return s
	}

	padding := width - currentWidth

	switch align {
	case "right":
		return spaces(padding) + s
	case "center":
		leftPad := padding / 2
		rightPad := padding - leftPad
		return spaces(leftPad) + s + spaces(rightPad)
	default: // left
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
