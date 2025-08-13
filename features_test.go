package termhyo

import (
	"bytes"
	"testing"
)

// TestIndividualFeatures tests specific features in isolation.
func TestIndividualFeatures(t *testing.T) {
	t.Run("HeaderStyleApply", func(t *testing.T) {
		style := HeaderStyle{
			Bold:            true,
			ForegroundColor: AnsiRed,
			BackgroundColor: AnsiBgYellow,
		}

		result := style.ApplyStyle("Test")
		expected := "\x1b[1m\x1b[31m\x1b[43mTest\x1b[0m"

		if result != expected {
			t.Errorf("HeaderStyle.ApplyStyle() = %q, expected %q", result, expected)
		}
	})

	t.Run("HeaderStyleIsEmpty", func(t *testing.T) {
		emptyStyle := HeaderStyle{}
		if !emptyStyle.isEmpty() {
			t.Error("Empty HeaderStyle should return true for isEmpty()")
		}

		nonEmptyStyle := HeaderStyle{Bold: true}
		if nonEmptyStyle.isEmpty() {
			t.Error("Non-empty HeaderStyle should return false for isEmpty()")
		}
	})

	t.Run("StringWidthWithEscapes", func(t *testing.T) {
		tests := []struct {
			input    string
			expected int
		}{
			{"hello", 5},
			{"\x1b[31mhello\x1b[0m", 5},
			{"こんにちは", 10},
			{"\x1b[31mこんにちは\x1b[0m", 10},
		}

		for _, test := range tests {
			result := StringWidth(test.input)
			if result != test.expected {
				t.Errorf("StringWidth(%q) = %d, expected %d", test.input, result, test.expected)
			}
		}
	})

	t.Run("BorderConfigDisabling", func(t *testing.T) {
		var buf bytes.Buffer
		columns := []Column{
			{Title: "Test", Width: 0, Align: "left"},
		}

		table := NewTable(&buf, columns)
		config := table.GetBorderConfig()
		config.Top = false
		config.Bottom = false
		table.SetBorderConfig(config)

		table.AddRow("Data")
		table.Render()

		output := buf.String()
		// Should not contain top/bottom border characters
		if bytes.Contains([]byte(output), []byte("┌")) || bytes.Contains([]byte(output), []byte("└")) {
			t.Error("Output should not contain top/bottom borders when disabled")
		}
	})
}

// TestConvenienceMethods tests the convenience methods for header styling.
func TestConvenienceMethods(t *testing.T) {
	tests := []struct {
		name    string
		setupFn func(*Table, HeaderStyle)
		checkFn func(*Table) bool
	}{
		{
			name: "SetHeaderStyleWithoutSeparator",
			setupFn: func(table *Table, style HeaderStyle) {
				table.SetHeaderStyleWithoutSeparator(style)
			},
			checkFn: func(table *Table) bool {
				config := table.GetBorderConfig()
				return !config.Middle
			},
		},
		{
			name: "SetHeaderStyleWithoutBorders",
			setupFn: func(table *Table, style HeaderStyle) {
				table.SetHeaderStyleWithoutBorders(style)
			},
			checkFn: func(table *Table) bool {
				config := table.GetBorderConfig()
				return !config.Top && !config.Middle && !config.Bottom
			},
		},
		{
			name: "SetHeaderStyleBorderless",
			setupFn: func(table *Table, style HeaderStyle) {
				table.SetHeaderStyleBorderless(style)
			},
			checkFn: func(table *Table) bool {
				config := table.GetBorderConfig()
				return !config.Top && !config.Middle && !config.Bottom &&
					!config.Left && !config.Right && config.Vertical
			},
		},
		{
			name: "SetHeaderStyleMinimal",
			setupFn: func(table *Table, style HeaderStyle) {
				table.SetHeaderStyleMinimal(style)
			},
			checkFn: func(table *Table) bool {
				config := table.GetBorderConfig()
				return !config.Top && !config.Middle && !config.Bottom &&
					!config.Left && !config.Right && !config.Vertical
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var buf bytes.Buffer
			columns := []Column{
				{Title: "Test", Width: 0, Align: "left"},
			}

			table := NewTable(&buf, columns)
			style := HeaderStyle{Bold: true}

			test.setupFn(table, style)

			if !test.checkFn(table) {
				t.Errorf("%s did not set expected border configuration", test.name)
			}

			// Check that header style was set
			if table.GetHeaderStyle().Bold != true {
				t.Errorf("%s did not set header style", test.name)
			}
		})
	}
}

// TestRenderModes tests different rendering modes.
func TestRenderModes(t *testing.T) {
	t.Run("BufferedMode", func(t *testing.T) {
		var buf bytes.Buffer
		columns := []Column{
			{Title: "Auto", Width: 0, Align: "left"}, // Auto width triggers buffered mode
		}

		table := NewTable(&buf, columns)
		table.AddRow("Test")
		table.Render()

		output := buf.String()
		if len(output) == 0 {
			t.Error("Buffered mode should produce output")
		}
	})

	t.Run("StreamingMode", func(t *testing.T) {
		var buf bytes.Buffer
		columns := []Column{
			{Title: "Fixed", Width: 10, Align: "left"}, // Fixed width triggers streaming mode
		}

		table := NewTable(&buf, columns)
		table.AddRow("Test")
		table.Render()

		output := buf.String()
		if len(output) == 0 {
			t.Error("Streaming mode should produce output")
		}
	})

	t.Run("NoAlignMode", func(t *testing.T) {
		var buf bytes.Buffer
		columns := []Column{
			{Title: "Test", Width: 0, Align: "left"},
		}

		table := NewTable(&buf, columns)
		table.SetAlign(false) // This should trigger streaming mode
		table.AddRow("Data")
		table.Render()

		output := buf.String()
		if len(output) == 0 {
			t.Error("No-align mode should produce output")
		}
	})
}
