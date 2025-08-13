package termhyo

// Column defines column properties.
type Column struct {
	Title    string // Column header title
	Width    int    // Column width (0 = auto-width)
	MaxWidth int    // Maximum width for auto-width columns (0 = no limit)
	Align    string // Alignment: "left", "center", "right"
}

// Cell represents a table cell.
type Cell struct {
	Content string // Cell content
	Align   string // Cell-specific alignment override
}

// Row represents a table row.
type Row struct {
	Cells []Cell // Row cells
}
