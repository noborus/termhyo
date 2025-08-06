package main

import (
	"os"

	"github.com/noborus/termhyo"
)

func main() {
	// Define columns with different alignments to show the effect
	columns := []termhyo.Column{
		{Title: "ID", Width: 0, Align: "right"},
		{Title: "Product Name", Width: 0, Align: "left"},
		{Title: "Price", Width: 0, Align: "right"},
		{Title: "Rating", Width: 0, Align: "center"},
		{Title: "Status", Width: 0, Align: "center"},
		{Title: "Description", Width: 0, Align: "left"},
	}

	// Create table with Markdown style
	table := termhyo.NewTableWithStyle(os.Stdout, columns, termhyo.MarkdownStyle)

	// Add sample data with varying lengths to show alignment
	table.AddRow("1", "Short", "$9.99", "★★★★★", "✓", "Brief")
	table.AddRow("42", "Very Long Product Name", "$199.99", "★★★☆☆", "✗", "This is a much longer description text")
	table.AddRow("123", "Medium Product", "$49.50", "★★★★☆", "✓", "Standard description")
	table.AddRow("7", "X", "$1.00", "★☆☆☆☆", "?", "Minimal")
	table.AddRow("9999", "Another Product Name", "$999.99", "★★★★★", "✓", "Great product with features")

	// Render the table
	table.Render()
}
