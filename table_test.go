package termhyo

import (
	"bytes"
	"flag"
	"os"
	"path/filepath"
	"testing"
)

var update = flag.Bool("update", false, "update golden files")

// TestTableOutput tests various table configurations against golden files.
func TestTableOutput(t *testing.T) {
	tests := []struct {
		name string
		fn   func() string
	}{
		{
			name: "basic_table",
			fn:   testBasicTable,
		},
		{
			name: "header_style_bold",
			fn:   testHeaderStyleBold,
		},
		{
			name: "header_style_background",
			fn:   testHeaderStyleBackground,
		},
		{
			name: "header_style_without_separator",
			fn:   testHeaderStyleWithoutSeparator,
		},
		{
			name: "header_style_without_borders",
			fn:   testHeaderStyleWithoutBorders,
		},
		{
			name: "header_style_borderless",
			fn:   testHeaderStyleBorderless,
		},
		{
			name: "header_style_minimal",
			fn:   testHeaderStyleMinimal,
		},
		{
			name: "different_border_styles",
			fn:   testDifferentBorderStyles,
		},
		{
			name: "markdown_table",
			fn:   testMarkdownTable,
		},
		{
			name: "markdown_with_header_style",
			fn:   testMarkdownWithHeaderStyle,
		},
		{
			name: "streaming_mode",
			fn:   testStreamingMode,
		},
		{
			name: "japanese_characters",
			fn:   testJapaneseCharacters,
		},
		{
			name: "escape_sequences",
			fn:   testEscapeSequences,
		},
		{
			name: "custom_borders",
			fn:   testCustomBorders,
		},
		{
			name: "no_align_mode",
			fn:   testNoAlignMode,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := tt.fn()
			goldenFile := filepath.Join("testdata", tt.name+".golden")

			if *update {
				// Create testdata directory if it doesn't exist
				if err := os.MkdirAll("testdata", 0o755); err != nil {
					t.Fatalf("Failed to create testdata directory: %v", err)
				}
				// Write the output to golden file
				if err := os.WriteFile(goldenFile, []byte(output), 0o644); err != nil {
					t.Fatalf("Failed to write golden file: %v", err)
				}
				return
			}

			// Read expected output from golden file
			expected, err := os.ReadFile(goldenFile)
			if err != nil {
				t.Fatalf("Failed to read golden file %s: %v", goldenFile, err)
			}

			if output != string(expected) {
				t.Errorf("Output mismatch for %s\nGot:\n%s\nExpected:\n%s", tt.name, output, string(expected))
			}
		})
	}
}

func testBasicTable() string {
	var buf bytes.Buffer
	columns := []Column{
		{Title: "ID", Width: 0, Align: Right},
		{Title: "Name", Width: 0, Align: Left},
		{Title: "Age", Width: 0, Align: Center},
		{Title: "City", Width: 0, Align: Left},
	}

	table := NewTable(&buf, columns)
	table.AddRow("1", "Alice", "25", "Tokyo")
	table.AddRow("2", "Bob", "30", "Osaka")
	table.AddRow("3", "Charlie", "35", "Kyoto")
	table.Render()

	return buf.String()
}

func testHeaderStyleBold() string {
	var buf bytes.Buffer
	columns := []Column{
		{Title: "Product", Width: 0, Align: Left},
		{Title: "Price", Width: 0, Align: Right},
		{Title: "Status", Width: 0, Align: Center},
	}

	table := NewTable(&buf, columns)
	table.SetHeaderStyle(BoldHeaderStyle())
	table.AddRow("Laptop", "$999.99", "Available")
	table.AddRow("Mouse", "$29.99", "Sold Out")
	table.Render()

	return buf.String()
}

func testHeaderStyleBackground() string {
	var buf bytes.Buffer
	columns := []Column{
		{Title: "Product", Width: 0, Align: Left},
		{Title: "Price", Width: 0, Align: Right},
		{Title: "Status", Width: 0, Align: Center},
	}

	table := NewTable(&buf, columns)
	headerStyle := HeaderStyle{
		Bold:            true,
		ForegroundColor: AnsiWhite,
		BackgroundColor: AnsiBgBlue,
	}
	table.SetHeaderStyle(headerStyle)
	table.AddRow("Laptop", "$999.99", "Available")
	table.AddRow("Mouse", "$29.99", "Sold Out")
	table.Render()

	return buf.String()
}

func testHeaderStyleWithoutSeparator() string {
	var buf bytes.Buffer
	columns := []Column{
		{Title: "Product", Width: 0, Align: Left},
		{Title: "Price", Width: 0, Align: Right},
		{Title: "Status", Width: 0, Align: Center},
	}

	table := NewTable(&buf, columns)
	headerStyle := HeaderStyle{
		Bold:            true,
		ForegroundColor: AnsiWhite,
		BackgroundColor: AnsiBgGreen,
	}
	table.SetHeaderStyleWithoutSeparator(headerStyle)
	table.AddRow("Laptop", "$999.99", "Available")
	table.AddRow("Mouse", "$29.99", "Sold Out")
	table.Render()

	return buf.String()
}

func testHeaderStyleWithoutBorders() string {
	var buf bytes.Buffer
	columns := []Column{
		{Title: "Product", Width: 0, Align: Left},
		{Title: "Price", Width: 0, Align: Right},
		{Title: "Status", Width: 0, Align: Center},
	}

	table := NewTable(&buf, columns)
	headerStyle := HeaderStyle{
		Bold:            true,
		ForegroundColor: AnsiWhite,
		BackgroundColor: AnsiBgRed,
	}
	table.SetHeaderStyleWithoutBorders(headerStyle)
	table.AddRow("Laptop", "$999.99", "Available")
	table.AddRow("Mouse", "$29.99", "Sold Out")
	table.Render()

	return buf.String()
}

func testHeaderStyleBorderless() string {
	var buf bytes.Buffer
	columns := []Column{
		{Title: "Product", Width: 0, Align: Left},
		{Title: "Price", Width: 0, Align: Right},
		{Title: "Status", Width: 0, Align: Center},
	}

	table := NewTable(&buf, columns)
	headerStyle := HeaderStyle{
		Bold:            true,
		ForegroundColor: AnsiWhite,
		BackgroundColor: AnsiBgCyan,
	}
	table.SetHeaderStyleBorderless(headerStyle)
	table.AddRow("Laptop", "$999.99", "Available")
	table.AddRow("Mouse", "$29.99", "Sold Out")
	table.Render()

	return buf.String()
}

func testHeaderStyleMinimal() string {
	var buf bytes.Buffer
	columns := []Column{
		{Title: "Product", Width: 0, Align: Left},
		{Title: "Price", Width: 0, Align: Right},
		{Title: "Status", Width: 0, Align: Center},
	}

	table := NewTable(&buf, columns)
	headerStyle := HeaderStyle{
		Bold:            true,
		ForegroundColor: AnsiWhite,
		BackgroundColor: AnsiBgMagenta,
	}
	table.SetHeaderStyleMinimal(headerStyle)
	table.AddRow("Laptop", "$999.99", "Available")
	table.AddRow("Mouse", "$29.99", "Sold Out")
	table.Render()

	return buf.String()
}

func testDifferentBorderStyles() string {
	var buf bytes.Buffer
	columns := []Column{
		{Title: "Style", Width: 0, Align: Left},
		{Title: "Description", Width: 0, Align: Left},
	}

	// Box Drawing Style
	buf.WriteString("=== Box Drawing Style ===\n")
	table1 := NewTable(&buf, columns, Border(BoxDrawingStyle))
	table1.AddRow("Box", "Unicode box drawing")
	table1.Render()
	buf.WriteString("\n")

	// ASCII Style
	buf.WriteString("=== ASCII Style ===\n")
	table2 := NewTable(&buf, columns, Border(ASCIIStyle))
	table2.AddRow("ASCII", "ASCII characters")
	table2.Render()
	buf.WriteString("\n")

	// Rounded Style
	buf.WriteString("=== Rounded Style ===\n")
	table3 := NewTable(&buf, columns, Border(RoundedStyle))
	table3.AddRow("Rounded", "Rounded corners")
	table3.Render()

	return buf.String()
}

func testMarkdownTable() string {
	var buf bytes.Buffer
	columns := []Column{
		{Title: "Feature", Width: 0, Align: Left},
		{Title: "Status", Width: 0, Align: Center},
		{Title: "Priority", Width: 0, Align: Right},
	}

	table := NewTable(&buf, columns, Border(MarkdownStyle))
	table.AddRow("Header styles", "Done", "High")
	table.AddRow("Border controls", "Done", "High")
	table.AddRow("Documentation", "In Progress", "Medium")
	table.Render()

	return buf.String()
}

func testMarkdownWithHeaderStyle() string {
	var buf bytes.Buffer
	columns := []Column{
		{Title: "Feature", Width: 0, Align: Left},
		{Title: "Status", Width: 0, Align: Center},
		{Title: "Priority", Width: 0, Align: Right},
	}

	headerStyle := HeaderStyle{
		Bold:            true,
		ForegroundColor: AnsiWhite,
		BackgroundColor: AnsiBgBlue,
	}
	table := NewTable(&buf, columns, Border(MarkdownStyle), Header(headerStyle))
	table.AddRow("Header styles", "Done", "High")
	table.AddRow("Border controls", "Done", "High")
	table.Render()

	return buf.String()
}

func testStreamingMode() string {
	var buf bytes.Buffer
	columns := []Column{
		{Title: "ID", Width: 3, Align: Right},       // Fixed width for streaming
		{Title: "Name", Width: 8, Align: Left},      // Fixed width for streaming
		{Title: "Status", Width: 10, Align: Center}, // Fixed width for streaming
	}

	table := NewTable(&buf, columns)
	table.AddRow("1", "Alice", "Active")
	table.AddRow("2", "Bob", "Inactive")
	table.Render()

	return buf.String()
}

func testJapaneseCharacters() string {
	var buf bytes.Buffer
	columns := []Column{
		{Title: "名前", Width: 0, Align: Left},
		{Title: "年齢", Width: 0, Align: Center},
		{Title: "職業", Width: 0, Align: Left},
	}

	table := NewTable(&buf, columns)
	table.AddRow("田中太郎", "30", "エンジニア")
	table.AddRow("佐藤花子", "25", "デザイナー")
	table.AddRow("鈴木一郎", "35", "マネージャー")
	table.Render()

	return buf.String()
}

func testEscapeSequences() string {
	var buf bytes.Buffer
	columns := []Column{
		{Title: "Type", Width: 0, Align: Left},
		{Title: "Content", Width: 0, Align: Left},
		{Title: "Width", Width: 0, Align: Right},
	}

	table := NewTable(&buf, columns)
	table.AddRow("Plain", "Hello World", "11")
	table.AddRow("Colored", "\x1b[31mRed Text\x1b[0m", "8")
	table.AddRow("Bold", "\x1b[1mBold Text\x1b[0m", "9")
	table.AddRow("Complex", "\x1b[1;32mBold Green\x1b[0m", "10")
	table.Render()

	return buf.String()
}

func testCustomBorders() string {
	var buf bytes.Buffer
	columns := []Column{
		{Title: "Column1", Width: 0, Align: Left},
		{Title: "Column2", Width: 0, Align: Center},
		{Title: "Column3", Width: 0, Align: Right},
	}

	table := NewTable(&buf, columns)

	// Custom border config - only internal vertical separators
	customConfig := BorderConfig{
		Chars: map[string]string{
			"vertical": " | ",
		},
		Top:      false,
		Bottom:   false,
		Middle:   false,
		Left:     false,
		Right:    false,
		Vertical: true,
		Padding:  true,
	}

	table.SetBorderConfig(customConfig)
	table.AddRow("Left", "Center", "Right")
	table.AddRow("Data1", "Data2", "Data3")
	table.Render()

	return buf.String()
}

func testNoAlignMode() string {
	var buf bytes.Buffer
	columns := []Column{
		{Title: "Raw1", Width: 0, Align: Left},
		{Title: "Raw2", Width: 0, Align: Center},
		{Title: "Raw3", Width: 0, Align: Right},
	}

	table := NewTable(&buf, columns)
	table.SetAlign(false) // Disable alignment
	table.AddRow("A", "B", "C")
	table.AddRow("LongText", "X", "Y")
	table.Render()

	return buf.String()
}
