package main

import (
	"fmt"
	"os"

	"github.com/noborus/termhyo"
)

func main() {
	// Define columns
	columns := []termhyo.Column{
		{Title: "ID", Width: 0, Align: termhyo.AlignRight},
		{Title: "Product", Width: 0, Align: termhyo.AlignLeft},
		{Title: "Price", Width: 0, Align: termhyo.AlignRight},
		{Title: "Status", Width: 0, Align: termhyo.AlignCenter},
	}

	fmt.Println("=== Default (No Header Style) ===")
	table1 := termhyo.NewTable(os.Stdout, columns)
	table1.AddRow("1", "Laptop", "$999.99", "Available")
	table1.AddRow("2", "Mouse", "$29.99", "Sold Out")
	table1.Render()
	fmt.Println()

	fmt.Println("=== Blue Background Header (Full Line) ===")
	table2 := termhyo.NewTable(os.Stdout, columns)
	blueHeaderStyle := termhyo.HeaderStyle{
		BackgroundColor: termhyo.AnsiBgBlue,
		ForegroundColor: termhyo.AnsiWhite,
		Bold:            true,
	}
	table2.SetHeaderStyle(blueHeaderStyle)
	table2.AddRow("1", "Laptop", "$999.99", "Available")
	table2.AddRow("2", "Mouse", "$29.99", "Sold Out")
	table2.Render()
	fmt.Println()

	fmt.Println("=== Green Background Header (Full Line) ===")
	table3 := termhyo.NewTable(os.Stdout, columns)
	greenHeaderStyle := termhyo.HeaderStyle{
		BackgroundColor: termhyo.AnsiBgGreen,
		ForegroundColor: termhyo.AnsiBlack,
		Bold:            true,
	}
	table3.SetHeaderStyle(greenHeaderStyle)
	table3.AddRow("1", "Laptop", "$999.99", "Available")
	table3.AddRow("2", "Mouse", "$29.99", "Sold Out")
	table3.Render()
	fmt.Println()

	fmt.Println("=== Yellow Background with Underline (Full Line) ===")
	table4 := termhyo.NewTable(os.Stdout, columns)
	yellowHeaderStyle := termhyo.HeaderStyle{
		BackgroundColor: termhyo.AnsiBgYellow,
		ForegroundColor: termhyo.AnsiBlack,
		Bold:            true,
		Underline:       true,
	}
	table4.SetHeaderStyle(yellowHeaderStyle)
	table4.AddRow("1", "Laptop", "$999.99", "Available")
	table4.AddRow("2", "Mouse", "$29.99", "Sold Out")
	table4.Render()
	fmt.Println()

	fmt.Println("=== True Color Header (Full Line) ===")
	table5 := termhyo.NewTable(os.Stdout, columns)
	trueColorStyle := termhyo.HeaderStyle{
		BackgroundColor: termhyo.TrueColorBg(75, 0, 130),    // Indigo background
		ForegroundColor: termhyo.TrueColorFg(255, 255, 255), // White text
		Bold:            true,
	}
	table5.SetHeaderStyle(trueColorStyle)
	table5.AddRow("1", "Laptop", "$999.99", "Available")
	table5.AddRow("2", "Mouse", "$29.99", "Sold Out")
	table5.Render()
	fmt.Println()
}
