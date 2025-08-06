package main

import (
	"os"

	"github.com/noborus/termhyo"
)

func main() {
	// Define columns
	columns := []termhyo.Column{
		{Title: "ID", Width: 0, Align: "right"},
		{Title: "名前", Width: 0, Align: "left"},
		{Title: "スコア", Width: 0, Align: "center"},
		{Title: "グレード", Width: 0, Align: "center"},
		{Title: "コメント", Width: 0, Align: "left"},
	}

	// Create table with default style
	table := termhyo.NewTable(os.Stdout, columns)

	// Add sample data with Japanese characters
	table.AddRow("1", "田中太郎", "95", "A", "素晴らしい成績です")
	table.AddRow("2", "佐藤花子", "87", "B", "良好な結果")
	table.AddRow("3", "山田三郎", "92", "A", "優秀です")
	table.AddRow("4", "鈴木一郎", "88", "B", "頑張りました")
	table.AddRow("5", "高橋美咲", "91", "A", "とても良い")

	// Render the table
	table.Render()
}
