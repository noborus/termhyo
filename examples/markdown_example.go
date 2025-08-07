package main

import (
	"os"

	"github.com/noborus/termhyo"
)

func main() {
	// Define columns
	columns := []termhyo.Column{
		{Title: "ID", Width: 0, Align: "right"},
		{Title: "Name", Width: 0, Align: "left"},
		{Title: "Score", Width: 0, Align: "center"},
		{Title: "Grade", Width: 0, Align: "center"},
	}

	// Create table with Markdown style
	table := termhyo.NewTable(os.Stdout, columns)
	table.SetBorderStyle(termhyo.MarkdownStyle)

	// Add sample data
	table.AddRow("1", "Alice Johnson", "95", "A")
	table.AddRow("2", "Bob Smith", "87", "B")
	table.AddRow("3", "Charlie Brown", "92", "A")
	table.AddRow("4", "Diana Prince", "88", "B")
	table.AddRow("5", "Edward Norton", "91", "A")

	// Render the table
	table.Render()
}
