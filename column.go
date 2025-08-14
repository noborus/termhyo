package termhyo

// Alignment represents text alignment options.
type Alignment string

const (
	// Default represents default/unspecified alignment.
	Default Alignment = ""
	// Left aligns text to the left.
	Left Alignment = "left"
	// Center aligns text to the center.
	Center Alignment = "center"
	// Right aligns text to the right.
	Right Alignment = "right"
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
	Align    Alignment // Alignment: Left, Center, Right
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
