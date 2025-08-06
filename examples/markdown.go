package main

import (
	"os"

	"github.com/noborus/termhyo"
)

func main() {
	// Define columns with alignment
	columns := []termhyo.Column{
		{Title: "ID", Width: 0, Align: "right"},
		{Title: "Name", Width: 0, Align: "left"},
		{Title: "Score", Width: 0, Align: "center"},
		{Title: "Grade", Width: 0, Align: "center"},
		{Title: "Comment", Width: 0, Align: "left"},
	}

	// Create table with Markdown style
	table := termhyo.NewTableWithStyle(os.Stdout, columns, termhyo.MarkdownStyle)

	// Add sample data
	table.AddRow("1", "Alice Johnson", "95", "A", "Excellent work")
	table.AddRow("2", "Bob Smith", "87", "B", "Good performance")
	table.AddRow("3", "Charlie Brown", "92", "A", "Very good")
	table.AddRow("4", "Diana Prince", "88", "B", "Solid effort")
	table.AddRow("5", "Edward Norton", "91", "A", "Outstanding")

	// Render the table
	table.Render()
}
