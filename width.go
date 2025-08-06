package termhyo

import (
	"strings"
	"unicode"

	"github.com/mattn/go-runewidth"
	"golang.org/x/text/unicode/norm"
)

// stringWidth returns the display width of a string
// This properly handles multibyte characters, combining characters, emojis, etc.
func stringWidth(s string) int {
	// Normalize the string to handle combining characters properly
	normalized := norm.NFC.String(s)
	return runewidth.StringWidth(normalized)
}

// truncateString truncates a string to fit within the specified display width
func truncateString(s string, maxWidth int) string {
	if maxWidth <= 0 {
		return ""
	}

	// Normalize the string first
	normalized := norm.NFC.String(s)

	if runewidth.StringWidth(normalized) <= maxWidth {
		return normalized
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

	w := 0
	var result []rune

	for _, r := range normalized {
		// Skip combining characters when calculating width
		// They should be included but don't add to display width
		if unicode.Is(unicode.Mn, r) || unicode.Is(unicode.Me, r) || unicode.Is(unicode.Mc, r) {
			result = append(result, r)
			continue
		}

		charWidth := runewidth.RuneWidth(r)

		if w+charWidth > targetWidth {
			break
		}

		result = append(result, r)
		w += charWidth
	}

	return string(result) + "..."
}

// padString pads a string to the specified display width with spaces
func padString(s string, width int, align string) string {
	currentWidth := stringWidth(s)
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

// spaces returns a string with n spaces
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
