package main

import (
	"fmt"
	"os"

	"github.com/noborus/termhyo"
)

func main() {
	// Define columns
	columns := []termhyo.Column{
		{Title: "Style", Width: 15, Align: termhyo.Left},
		{Title: "Description", Width: 30, Align: termhyo.Left},
		{Title: "Unicode", Width: 10, Align: termhyo.Center},
	}

	// Sample data
	data := [][]string{
		{"BoxDrawing", "Unicode box drawing (default)", "Yes"},
		{"ASCII", "Simple ASCII characters", "No"},
		{"Rounded", "Rounded corners", "Yes"},
		{"Double", "Double line borders", "Yes"},
		{"Minimal", "Minimal border style", "No"},
	}

	styles := []termhyo.BorderStyle{
		termhyo.BoxDrawingStyle,
		termhyo.ASCIIStyle,
		termhyo.RoundedStyle,
		termhyo.DoubleStyle,
		termhyo.MinimalStyle,
	}

	// Demonstrate each border style
	for i, style := range styles {
		fmt.Printf("\n=== %s ===\n", data[i][0])

		table := termhyo.NewTable(os.Stdout, columns, termhyo.Border(style))
		for _, row := range data {
			table.AddRow(row...)
		}
		table.Render()
	}
}
