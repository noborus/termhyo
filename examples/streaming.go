package main

import (
	"os"
	"time"

	"github.com/noborus/termhyo"
)

func main() {
	// Define columns with fixed width for streaming mode
	columns := []termhyo.Column{
		{Title: "Time", Width: 8, Align: ""},
		{Title: "Event", Width: 20, Align: ""},
		{Title: "Status", Width: 10, Align: ""},
	}

	// Create table with fixed width (enables streaming mode)
	table := termhyo.NewTable(os.Stdout, columns)

	// Simulate real-time data streaming
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
		time.Sleep(500 * time.Millisecond)
	}
}
