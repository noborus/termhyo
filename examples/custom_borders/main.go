package main

import (
	"fmt"
	"os"

	"github.com/noborus/termhyo"
)

func main() {
	// Define columns
	columns := []termhyo.Column{
		{Title: "Style", Width: 0, Align: termhyo.Left},
		{Title: "Description", Width: 0, Align: termhyo.Left},
		{Title: "Borders", Width: 0, Align: termhyo.Center},
	}

	// Test different border configurations
	testConfigs := []struct {
		name   string
		style  termhyo.BorderStyle
		desc   string
		border string
	}{
		{"Box Drawing", termhyo.BoxDrawingStyle, "Full borders with Unicode", "All"},
		{"ASCII", termhyo.ASCIIStyle, "Full borders with ASCII", "All"},
		{"Minimal", termhyo.MinimalStyle, "No visible borders", "None"},
		{"Markdown", termhyo.MarkdownStyle, "Markdown table format", "Middle only"},
		{"TSV", termhyo.TSVStyle, "Tab-separated values", "Tabs only"},
	}

	for _, config := range testConfigs {
		fmt.Printf("=== %s Style ===\n", config.name)
		table := termhyo.NewTable(os.Stdout, columns, termhyo.Border(config.style))

		// Add header
		table.AddRow("Test", config.desc, config.border)
		table.AddRow("Data", "Sample row", "Display")

		table.Render()

		// Add spacing between different styles
		fmt.Printf("\n")
	}

	// Example of custom border configuration - only internal separators
	fmt.Printf("=== Custom: Internal separators only ===\n")
	customColumns := []termhyo.Column{
		{Title: "Column1", Width: 0, Align: termhyo.Left},
		{Title: "Column2", Width: 0, Align: termhyo.Center},
		{Title: "Column3", Width: 0, Align: termhyo.Right},
	}

	customTable := termhyo.NewTable(os.Stdout, customColumns)

	// Create custom border config - only internal vertical separators
	customConfig := termhyo.BorderConfig{
		Chars: map[string]string{
			"vertical": " | ",
		},
		Top:      false,
		Bottom:   false,
		Middle:   false,
		Left:     false,
		Right:    false,
		Vertical: true,
	}

	customTable.SetBorderConfig(customConfig)
	customTable.AddRow("Left", "Center", "Right")
	customTable.AddRow("Data1", "Data2", "Data3")
	customTable.Render()
}
