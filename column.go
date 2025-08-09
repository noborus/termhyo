package termhyo

// Column defines column properties
type Column struct {
	Key      string
	Title    string
	Width    int    // 0 means auto-width
	MaxWidth int    // maximum width for auto-width columns
	Align    string // "left", "right", "center"
}

// Cell represents a table cell
type Cell struct {
	Content string
	Align   string // overrides column alignment if set
}

// Row represents a table row
type Row struct {
	Cells []Cell
}
