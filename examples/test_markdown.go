package main

import (
	"bytes"
	"fmt"

	"github.com/noborus/termhyo"
)

func main() {
	// Create a buffer to capture output
	var buf bytes.Buffer

	// Define simple columns
	columns := []termhyo.Column{
		{Title: "ID", Width: 0, Align: "right"},
		{Title: "Name", Width: 0, Align: "left"},
		{Title: "Score", Width: 0, Align: "center"},
	}

	// Create table with Markdown style
	table := termhyo.NewTableWithStyle(&buf, columns, termhyo.MarkdownStyle)

	// Add test data
	table.AddRow("1", "Alice", "95")
	table.AddRow("42", "Bob", "87")

	// Render the table
	err := table.Render()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Print the result
	fmt.Print("Generated Markdown table:\n")
	fmt.Print(buf.String())
}
