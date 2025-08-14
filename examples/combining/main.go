package main

import (
	"os"

	"github.com/noborus/termhyo"
)

func main() {
	// Define columns
	columns := []termhyo.Column{
		{Title: "Text Type", Width: 0, Align: termhyo.Left},
		{Title: "Example", Width: 0, Align: termhyo.Left},
		{Title: "Description", Width: 0, Align: termhyo.Left},
	}

	// Create table with default style
	table := termhyo.NewTable(os.Stdout, columns)

	// Add sample data with combining characters and complex Unicode
	table.AddRow("Combining", "é (e + ́)", "e with combining acute accent")
	table.AddRow("Combining", "ñ (n + ̃)", "n with combining tilde")
	table.AddRow("Combining", "ö (o + ̈)", "o with combining diaeresis")
	table.AddRow("Emoji ZWJ", "👨‍👩‍👧‍👦", "Family emoji with ZWJ sequences")
	table.AddRow("Emoji Mod", "👋🏻", "Waving hand with skin tone modifier")
	table.AddRow("Emoji Var", "👍️", "Thumbs up with variation selector")
	table.AddRow("Complex", "🏳️‍🌈", "Rainbow flag with ZWJ sequence")
	table.AddRow("Arabic", "مُرَكَّب", "Arabic text with diacritics")

	// Render the table
	table.Render()
}
