package termhyo

// Alignment represents text alignment options.
type Alignment string

const (
	// AlignDefault represents default/unspecified alignment.
	AlignDefault Alignment = ""
	// AlignLeft aligns text to the left.
	AlignLeft Alignment = "left"
	// AlignCenter aligns text to the center.
	AlignCenter Alignment = "center"
	// AlignRight aligns text to the right.
	AlignRight Alignment = "right"
)

// String returns the string representation of the alignment.
func (a Alignment) String() string {
	return string(a)
}

// Column defines column properties.
type Column struct {
	Title    string    // Column header title
	Width    int       // Column width (0 = auto-width)
	MaxWidth int       // Maximum width for auto-width columns (0 = no limit)
	Align    Alignment // Alignment: AlignLeft, AlignCenter, AlignRight
}

// Cell represents a table cell.
type Cell struct {
	Content string    // Cell content
	Align   Alignment // Cell-specific alignment override
}

// Row represents a table row.
type Row struct {
	Cells []Cell // Row cells
}
