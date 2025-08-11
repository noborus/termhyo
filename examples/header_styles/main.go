package main

import (
	"fmt"
	"os"

	"github.com/noborus/termhyo"
)

func main() {
	// Define columns
	columns := []termhyo.Column{
		{Title: "ID", Width: 0, Align: "right"},
		{Title: "Product Name", Width: 0, Align: "left"},
		{Title: "Price", Width: 0, Align: "right"},
		{Title: "Status", Width: 0, Align: "center"},
	}

	fmt.Println("=== Default Header Style (with separator line) ===")
	table1 := termhyo.NewTable(os.Stdout, columns)
	table1.AddRow("1", "Laptop", "$999.99", "Available")
	table1.AddRow("2", "Mouse", "$29.99", "Sold Out")
	table1.Render()
	fmt.Println()

	fmt.Println("=== Blue Background Headers (no separator line needed) ===")
	table2 := termhyo.NewTable(os.Stdout, columns)
	blueHeaderStyle := termhyo.HeaderStyle{
		Bold:            true,
		ForegroundColor: termhyo.AnsiWhite,
		BackgroundColor: termhyo.AnsiBgBlue,
	}
	table2.SetHeaderStyle(blueHeaderStyle)
	// Disable the header separator line and top/bottom borders for cleaner look
	borderConfig := table2.GetBorderConfig()
	borderConfig.Middle = false
	borderConfig.Top = false
	borderConfig.Bottom = false
	table2.SetBorderConfig(borderConfig)
	table2.AddRow("1", "Laptop", "$999.99", "Available")
	table2.AddRow("2", "Mouse", "$29.99", "Sold Out")
	table2.Render()
	fmt.Println()

	fmt.Println("=== Green Background Headers (no separator line) ===")
	table3 := termhyo.NewTable(os.Stdout, columns)
	greenHeaderStyle := termhyo.HeaderStyle{
		Bold:            true,
		ForegroundColor: termhyo.AnsiWhite,
		BackgroundColor: termhyo.AnsiBgGreen,
	}
	table3.SetHeaderStyle(greenHeaderStyle)
	// Disable borders for clean look
	borderConfig3 := table3.GetBorderConfig()
	borderConfig3.Middle = false
	borderConfig3.Top = false
	borderConfig3.Bottom = false
	table3.SetBorderConfig(borderConfig3)
	table3.AddRow("1", "Laptop", "$999.99", "Available")
	table3.AddRow("2", "Mouse", "$29.99", "Sold Out")
	table3.Render()
	fmt.Println()

	fmt.Println("=== Yellow Background with Bold Text (no separator) ===")
	table4 := termhyo.NewTable(os.Stdout, columns)
	yellowHeaderStyle := termhyo.HeaderStyle{
		Bold:            true,
		ForegroundColor: termhyo.AnsiBlack,
		BackgroundColor: termhyo.AnsiBgYellow,
	}
	table4.SetHeaderStyle(yellowHeaderStyle)
	// Disable borders for clean look
	borderConfig4 := table4.GetBorderConfig()
	borderConfig4.Middle = false
	borderConfig4.Top = false
	borderConfig4.Bottom = false
	table4.SetBorderConfig(borderConfig4)
	table4.AddRow("1", "Laptop", "$999.99", "Available")
	table4.AddRow("2", "Mouse", "$29.99", "Sold Out")
	table4.Render()
	fmt.Println()

	fmt.Println("=== Cyan Background with Underline (no separator) ===")
	table5 := termhyo.NewTable(os.Stdout, columns)
	cyanHeaderStyle := termhyo.HeaderStyle{
		Bold:            true,
		Underline:       true,
		ForegroundColor: termhyo.AnsiBlack,
		BackgroundColor: termhyo.AnsiBgCyan,
	}
	table5.SetHeaderStyle(cyanHeaderStyle)
	// Disable borders for clean look
	borderConfig5 := table5.GetBorderConfig()
	borderConfig5.Middle = false
	borderConfig5.Top = false
	borderConfig5.Bottom = false
	table5.SetBorderConfig(borderConfig5)
	table5.AddRow("1", "Laptop", "$999.99", "Available")
	table5.AddRow("2", "Mouse", "$29.99", "Sold Out")
	table5.Render()
	fmt.Println()

	fmt.Println("=== True Color Header (no separator) ===")
	table6 := termhyo.NewTable(os.Stdout, columns)
	trueColorStyle := termhyo.HeaderStyle{
		Bold:            true,
		ForegroundColor: termhyo.TrueColorFg(255, 255, 255), // White text
		BackgroundColor: termhyo.TrueColorBg(75, 0, 130),    // Indigo background
	}
	table6.SetHeaderStyle(trueColorStyle)
	// Disable borders for clean look
	borderConfig6 := table6.GetBorderConfig()
	borderConfig6.Middle = false
	borderConfig6.Top = false
	borderConfig6.Bottom = false
	table6.SetBorderConfig(borderConfig6)
	table6.AddRow("1", "Laptop", "$999.99", "Available")
	table6.AddRow("2", "Mouse", "$29.99", "Sold Out")
	table6.Render()
	fmt.Println()

	fmt.Println("=== Reverse Video Header (no separator) ===")
	table7 := termhyo.NewTable(os.Stdout, columns)
	reverseStyle := termhyo.HeaderStyle{
		Bold:    true,
		Reverse: true, // Swap foreground and background colors
	}
	table7.SetHeaderStyle(reverseStyle)
	// Disable borders for clean look
	borderConfig7 := table7.GetBorderConfig()
	borderConfig7.Middle = false
	borderConfig7.Top = false
	borderConfig7.Bottom = false
	table7.SetBorderConfig(borderConfig7)
	table7.AddRow("1", "Laptop", "$999.99", "Available")
	table7.AddRow("2", "Mouse", "$29.99", "Sold Out")
	table7.Render()
	fmt.Println()

	fmt.Println("=== Comparison: Same style with and without separator ===")
	headerStyle := termhyo.HeaderStyle{
		Bold:            true,
		ForegroundColor: termhyo.AnsiWhite,
		BackgroundColor: termhyo.AnsiBgMagenta,
	}

	fmt.Println("With separator line:")
	tableWith := termhyo.NewTable(os.Stdout, columns)
	tableWith.SetHeaderStyle(headerStyle)
	tableWith.AddRow("1", "Laptop", "$999.99", "Available")
	tableWith.AddRow("2", "Mouse", "$29.99", "Sold Out")
	tableWith.Render()
	fmt.Println()

	fmt.Println("Without separator line (cleaner look):")
	tableWithout := termhyo.NewTable(os.Stdout, columns)
	tableWithout.SetHeaderStyle(headerStyle)
	borderConfigWithout := tableWithout.GetBorderConfig()
	borderConfigWithout.Middle = false
	borderConfigWithout.Top = false
	borderConfigWithout.Bottom = false
	tableWithout.SetBorderConfig(borderConfigWithout)
	tableWithout.AddRow("1", "Laptop", "$999.99", "Available")
	tableWithout.AddRow("2", "Mouse", "$29.99", "Sold Out")
	tableWithout.Render()
	fmt.Println()

	fmt.Println("=== Using convenience method: SetHeaderStyleWithoutSeparator ===")
	tableConvenience := termhyo.NewTable(os.Stdout, columns)
	redStyle := termhyo.HeaderStyle{
		Bold:            true,
		ForegroundColor: termhyo.AnsiWhite,
		BackgroundColor: termhyo.AnsiBgRed,
	}
	// This automatically disables the header separator line
	tableConvenience.SetHeaderStyleWithoutSeparator(redStyle)
	tableConvenience.AddRow("1", "Laptop", "$999.99", "Available")
	tableConvenience.AddRow("2", "Mouse", "$29.99", "Sold Out")
	tableConvenience.Render()
	fmt.Println()

	fmt.Println("=== Using convenience method: SetHeaderStyleWithoutBorders ===")
	tableClean := termhyo.NewTable(os.Stdout, columns)
	purpleStyle := termhyo.HeaderStyle{
		Bold:            true,
		ForegroundColor: termhyo.AnsiWhite,
		BackgroundColor: termhyo.AnsiBgMagenta,
	}
	// This automatically disables all horizontal borders for the cleanest look
	tableClean.SetHeaderStyleWithoutBorders(purpleStyle)
	tableClean.AddRow("1", "Laptop", "$999.99", "Available")
	tableClean.AddRow("2", "Mouse", "$29.99", "Sold Out")
	tableClean.Render()
	fmt.Println()

	fmt.Println("=== Completely borderless table (no left/right borders) ===")
	tableBorderless := termhyo.NewTable(os.Stdout, columns)
	orangeStyle := termhyo.HeaderStyle{
		Bold:            true,
		ForegroundColor: termhyo.AnsiBlack,
		BackgroundColor: termhyo.AnsiBgBrightYellow,
	}
	tableBorderless.SetHeaderStyle(orangeStyle)
	// Disable ALL borders for the most minimal look
	borderConfigBorderless := tableBorderless.GetBorderConfig()
	borderConfigBorderless.Top = false
	borderConfigBorderless.Bottom = false
	borderConfigBorderless.Middle = false
	borderConfigBorderless.Left = false
	borderConfigBorderless.Right = false
	// Keep internal vertical separators for column distinction
	borderConfigBorderless.Vertical = true
	tableBorderless.SetBorderConfig(borderConfigBorderless)
	tableBorderless.AddRow("1", "Laptop", "$999.99", "Available")
	tableBorderless.AddRow("2", "Mouse", "$29.99", "Sold Out")
	tableBorderless.Render()
	fmt.Println()

	fmt.Println("=== Completely clean table (no borders at all) ===")
	tableCleanest := termhyo.NewTable(os.Stdout, columns)
	tealStyle := termhyo.HeaderStyle{
		Bold:            true,
		ForegroundColor: termhyo.AnsiWhite,
		BackgroundColor: termhyo.AnsiBgCyan,
	}
	tableCleanest.SetHeaderStyle(tealStyle)
	// Disable absolutely ALL borders and separators
	borderConfigCleanest := tableCleanest.GetBorderConfig()
	borderConfigCleanest.Top = false
	borderConfigCleanest.Bottom = false
	borderConfigCleanest.Middle = false
	borderConfigCleanest.Left = false
	borderConfigCleanest.Right = false
	borderConfigCleanest.Vertical = false // No internal separators either
	tableCleanest.SetBorderConfig(borderConfigCleanest)
	tableCleanest.AddRow("1", "Laptop", "$999.99", "Available")
	tableCleanest.AddRow("2", "Mouse", "$29.99", "Sold Out")
	tableCleanest.Render()
	fmt.Println()

	fmt.Println("=== Using SetHeaderStyleBorderless (convenience method) ===")
	tableBorderlessConvenience := termhyo.NewTable(os.Stdout, columns)
	greenStyle := termhyo.HeaderStyle{
		Bold:            true,
		ForegroundColor: termhyo.AnsiWhite,
		BackgroundColor: termhyo.AnsiBgGreen,
	}
	// This automatically disables all borders but keeps column separators
	tableBorderlessConvenience.SetHeaderStyleBorderless(greenStyle)
	tableBorderlessConvenience.AddRow("1", "Laptop", "$999.99", "Available")
	tableBorderlessConvenience.AddRow("2", "Mouse", "$29.99", "Sold Out")
	tableBorderlessConvenience.Render()
	fmt.Println()

	fmt.Println("=== Using SetHeaderStyleMinimal (ultimate clean) ===")
	tableMinimal := termhyo.NewTable(os.Stdout, columns)
	blueStyle := termhyo.HeaderStyle{
		Bold:            true,
		ForegroundColor: termhyo.AnsiWhite,
		BackgroundColor: termhyo.AnsiBgBlue,
	}
	// This disables absolutely everything - only whitespace and header styling
	tableMinimal.SetHeaderStyleMinimal(blueStyle)
	tableMinimal.AddRow("1", "Laptop", "$999.99", "Available")
	tableMinimal.AddRow("2", "Mouse", "$29.99", "Sold Out")
	tableMinimal.Render()
}
