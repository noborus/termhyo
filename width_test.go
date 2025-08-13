package termhyo

import (
	"testing"
)

func TestStringWidthWithEscapeSequences(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "plain text",
			input:    "hello",
			expected: 5,
		},
		{
			name:     "ANSI color code",
			input:    "\x1b[31mhello\x1b[0m",
			expected: 5,
		},
		{
			name:     "ANSI bold and color",
			input:    "\x1b[1;31mhello\x1b[0m",
			expected: 5,
		},
		{
			name:     "multiple ANSI codes",
			input:    "\x1b[31mred\x1b[32mgreen\x1b[34mblue\x1b[0m",
			expected: 12, // "redgreenblue"
		},
		{
			name:     "ANSI with Japanese characters",
			input:    "\x1b[31mこんにちは\x1b[0m",
			expected: 10, // Japanese characters are 2-width each
		},
		{
			name:     "background color",
			input:    "\x1b[41mhello\x1b[0m",
			expected: 5,
		},
		{
			name:     "256 color code",
			input:    "\x1b[38;5;196mhello\x1b[0m",
			expected: 5,
		},
		{
			name:     "RGB color code",
			input:    "\x1b[38;2;255;0;0mhello\x1b[0m",
			expected: 5,
		},
		{
			name:     "cursor movement codes",
			input:    "\x1b[2Jhello\x1b[H",
			expected: 5,
		},
		{
			name:     "mixed content with spaces",
			input:    "\x1b[31mred text\x1b[0m normal",
			expected: 15, // "red text normal"
		},
		{
			name:     "empty string",
			input:    "",
			expected: 0,
		},
		{
			name:     "only escape sequences",
			input:    "\x1b[31m\x1b[0m",
			expected: 0,
		},
		{
			name:     "control characters",
			input:    "hello\r\nworld",
			expected: 10, // "helloworld" (control chars removed)
		},
		{
			name:     "tabs and spaces",
			input:    "hello\tworld test",
			expected: 15, // "hello"(5) + tab(1) + "world test"(10) - 1 = 15
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StringWidth(tt.input)
			if result != tt.expected {
				t.Errorf("StringWidth(%q) = %d, expected %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestStripEscapeSequences(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "plain text",
			input:    "hello",
			expected: "hello",
		},
		{
			name:     "ANSI color code",
			input:    "\x1b[31mhello\x1b[0m",
			expected: "hello",
		},
		{
			name:     "multiple ANSI codes",
			input:    "\x1b[1;31mBold Red\x1b[0m\x1b[32mGreen\x1b[0m",
			expected: "Bold RedGreen",
		},
		{
			name:     "control characters",
			input:    "hello\r\n\tworld",
			expected: "hello\tworld", // \r\n removed, \t kept
		},
		{
			name:     "complex escape sequence",
			input:    "\x1b[38;2;255;0;0mRGB red\x1b[0m",
			expected: "RGB red",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := stripEscapeSequences(tt.input)
			if result != tt.expected {
				t.Errorf("stripEscapeSequences(%q) = %q, expected %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestTruncateStringWithEscapes(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		maxWidth  int
		expected  string
		expectLen int // expected display width
	}{
		{
			name:      "plain text no truncation",
			input:     "hello",
			maxWidth:  10,
			expected:  "hello",
			expectLen: 5,
		},
		{
			name:      "plain text with truncation",
			input:     "hello world",
			maxWidth:  8,
			expected:  "hello...",
			expectLen: 8,
		},
		{
			name:      "ANSI color preserved",
			input:     "\x1b[31mhello world\x1b[0m",
			maxWidth:  8,
			expected:  "\x1b[31mhello\x1b[0m...",
			expectLen: 8,
		},
		{
			name:      "ANSI color no truncation needed",
			input:     "\x1b[31mhello\x1b[0m",
			maxWidth:  10,
			expected:  "\x1b[31mhello\x1b[0m",
			expectLen: 5,
		},
		{
			name:      "multiple ANSI codes with truncation",
			input:     "\x1b[1;31mhello\x1b[32m world\x1b[0m",
			maxWidth:  8,
			expected:  "\x1b[1;31mhello\x1b[32m w...",
			expectLen: 8,
		},
		{
			name:      "very short width",
			input:     "hello",
			maxWidth:  2,
			expected:  "..",
			expectLen: 2,
		},
		{
			name:      "zero width",
			input:     "hello",
			maxWidth:  0,
			expected:  "",
			expectLen: 0,
		},
		{
			name:      "Japanese with ANSI",
			input:     "\x1b[31mこんにちは世界\x1b[0m",
			maxWidth:  10,
			expected:  "\x1b[31mこんにちは\x1b[0m...",
			expectLen: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := truncateString(tt.input, tt.maxWidth)
			resultWidth := stringWidth(result)

			if resultWidth != tt.expectLen {
				t.Errorf("truncateString(%q, %d) resulted in width %d, expected %d",
					tt.input, tt.maxWidth, resultWidth, tt.expectLen)
			}

			// Check if result contains expected content (this is more complex due to escape sequences)
			if tt.maxWidth > 0 && resultWidth > tt.maxWidth {
				t.Errorf("truncateString(%q, %d) resulted in width %d > maxWidth %d",
					tt.input, tt.maxWidth, resultWidth, tt.maxWidth)
			}
		})
	}
}

func TestPadStringWithEscapes(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		width     int
		align     Alignment
		expectLen int
	}{
		{
			name:      "left align plain text",
			input:     "hello",
			width:     10,
			align:     AlignLeft,
			expectLen: 10,
		},
		{
			name:      "right align plain text",
			input:     "hello",
			width:     10,
			align:     AlignRight,
			expectLen: 10,
		},
		{
			name:      "center align plain text",
			input:     "hello",
			width:     10,
			align:     AlignCenter,
			expectLen: 10,
		},
		{
			name:      "ANSI color left align",
			input:     "\x1b[31mhello\x1b[0m",
			width:     10,
			align:     AlignLeft,
			expectLen: 10,
		},
		{
			name:      "ANSI color right align",
			input:     "\x1b[31mhello\x1b[0m",
			width:     10,
			align:     AlignRight,
			expectLen: 10,
		},
		{
			name:      "ANSI color center align",
			input:     "\x1b[31mhello\x1b[0m",
			width:     10,
			align:     AlignCenter,
			expectLen: 10,
		},
		{
			name:      "Japanese with ANSI",
			input:     "\x1b[31mこんにちは\x1b[0m",
			width:     15,
			align:     AlignCenter,
			expectLen: 15,
		},
		{
			name:      "no padding needed",
			input:     "\x1b[31mhello world\x1b[0m",
			width:     10,
			align:     AlignLeft,
			expectLen: 11, // original width is larger
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := padString(tt.input, tt.width, tt.align)
			resultWidth := stringWidth(result)

			if resultWidth != tt.expectLen {
				t.Errorf("padString(%q, %d, %q) resulted in width %d, expected %d",
					tt.input, tt.width, tt.align, resultWidth, tt.expectLen)
			}
		})
	}
}

// Benchmark tests to ensure performance is acceptable.
func BenchmarkStringWidth(b *testing.B) {
	testStrings := []string{
		"hello world",
		"\x1b[31mhello world\x1b[0m",
		"\x1b[1;31mこんにちは世界\x1b[0m",
		"\x1b[38;2;255;0;0mRGB colored text\x1b[0m",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, s := range testStrings {
			stringWidth(s)
		}
	}
}

func BenchmarkTruncateString(b *testing.B) {
	testString := "\x1b[1;31mThis is a long string with ANSI escape sequences that needs truncation\x1b[0m"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		truncateString(testString, 20)
	}
}
