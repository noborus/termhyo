package main

import (
	"os"

	"github.com/noborus/termhyo"
)

func main() {
	columns := []termhyo.Column{
		{Title: "ID", Width: 0, Align: termhyo.Right},
		{Title: "Name", Width: 0, Align: termhyo.Left},
		{Title: "Score", Width: 0, Align: termhyo.Center},
	}

	table := termhyo.NewTable(os.Stdout, columns, termhyo.Border(termhyo.VerticalBarStyle), termhyo.AutoAlign(false))
	table.AddRow("1", "Alice", "85")
	table.AddRow("2", "Bob", "92")
	table.AddRow("3", "Charlie", "78")
	table.Render()
}
