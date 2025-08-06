package main

import (
	"os"

	"github.com/noborus/termhyo"
)

func main() {
	// Define columns
	columns := []termhyo.Column{
		{Title: "Status", Width: 0, Align: "center"},
		{Title: "Task", Width: 0, Align: "left"},
		{Title: "Progress", Width: 0, Align: "center"},
		{Title: "Assignee", Width: 0, Align: "left"},
	}

	// Create table with default style
	table := termhyo.NewTable(os.Stdout, columns)

	// Add sample data with emojis and Unicode characters
	table.AddRow("✅", "完了したタスク", "100%", "田中 📧")
	table.AddRow("🔄", "進行中のタスク", "75%", "佐藤 🚀")
	table.AddRow("⏳", "待機中のタスク", "0%", "山田 ⭐")
	table.AddRow("❌", "失敗したタスク", "0%", "鈴木 💻")
	table.AddRow("🎯", "重要なタスク", "50%", "高橋 🌟")

	// Render the table
	table.Render()
}
