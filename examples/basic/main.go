package main

import (
	"os"

	"github.com/noborus/termhyo"
)

func main() {
	// Define columns
	columns := []termhyo.Column{
		{Title: "ID", Width: 0, Align: termhyo.Right},
		{Title: "Name", Width: 0, Align: termhyo.Left},
		{Title: "Score", Width: 0, Align: termhyo.Center},
		{Title: "Grade", Width: 0, Align: termhyo.Center},
	}

	// Create table with default style
	table := termhyo.NewTable(os.Stdout, columns)

	// Add sample data
	table.AddRow("1", "Alice Johnson", "95", "A")
	table.AddRow("2", "Bob Smith", "87", "B")
	table.AddRow("3", "Charlie Brown", "92", "A")
	table.AddRow("4", "Diana Prince", "88", "B")
	table.AddRow("5", "Edward Norton", "91", "A")

	// Render the table
	table.Render()
}
