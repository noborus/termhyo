package main

import (
	"fmt"
	"os"

	"github.com/noborus/termhyo"
)

func main() {
	fmt.Println("Testing Markdown table...")

	// Define columns
	columns := []termhyo.Column{
		{Title: "ID", Width: 0, Align: "right"},
		{Title: "Name", Width: 0, Align: "left"},
		{Title: "Score", Width: 0, Align: "center"},
	}

	// Create table with Markdown style
	table := termhyo.NewTableWithStyle(os.Stdout, columns, termhyo.MarkdownStyle)

	// Add test data
	table.AddRow("1", "Alice", "95")
	table.AddRow("42", "Bob", "87")

	// Render the table
	fmt.Println("Rendering table...")
	err := table.Render()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println("Done.")
}
