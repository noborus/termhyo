package main

import (
	"os"

	"github.com/noborus/termhyo"
)

func main() {
	// Define columns with fixed width for streaming Markdown
	columns := []termhyo.Column{
		{Title: "Time", Width: 10, Align: "left"},
		{Title: "Event", Width: 20, Align: "left"},
		{Title: "Status", Width: 8, Align: "center"},
	}

	// Create table with Markdown style
	table := termhyo.NewTableWithStyle(os.Stdout, columns, termhyo.MarkdownStyle)
	table.SetNoAlign(true) // Disable alignment for streaming mode
	// Simulate real-time data streaming in Markdown format
	events := [][]string{
		{"09:00:00", "System startup", "OK"},
		{"09:00:15", "Loading config", "OK"},
		{"09:00:30", "Database connect", "OK"},
		{"09:00:45", "Cache warming", "OK"},
		{"09:01:00", "Service ready", "OK"},
	}

	for _, event := range events {
		table.AddRow(event...)
		// Simulate delay between events
		//time.Sleep(500 * time.Millisecond)
	}

	// Complete the table
	table.Render()
}
