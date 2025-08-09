package main

import (
	"fmt"
	"os"

	"github.com/noborus/termhyo"
)

func main() {
	fmt.Println("Starting test...")

	// Define columns
	columns := []termhyo.Column{
		{Title: "Test", Width: 0, Align: "left"},
		{Title: "Value", Width: 0, Align: "left"},
	}

	fmt.Println("Creating table...")
	// Create table with default style
	table := termhyo.NewTable(os.Stdout, columns)

	fmt.Println("Adding simple rows...")
	// Add simple data first
	table.AddRow("Simple", "ASCII text")

	fmt.Println("Adding Japanese...")
	table.AddRow("Japanese", "こんにちは")

	fmt.Println("Rendering...")
	// Render the table
	table.Render()

	fmt.Println("Done!")
}
