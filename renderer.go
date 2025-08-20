package termhyo

// Renderer defines the interface for different rendering strategies.
type Renderer interface {
	AddRow(table *Table, row Row) error
	Render(table *Table) error
	IsRendered() bool
}

// RenderMode defines when to start rendering.
type RenderMode int

const (
	// BufferedMode: collect all rows before rendering (for auto-width)
	BufferedMode RenderMode = iota
	// StreamingMode: render immediately (for fixed-width)
	StreamingMode
)

// String returns the string representation of the RenderMode.
func (m RenderMode) String() string {
	switch m {
	case BufferedMode:
		return "Buffered"
	case StreamingMode:
		return "Streaming"
	default:
		return "Unknown"
	}
}

// Buffered implements buffered rendering strategy.
type Buffered struct {
	rendered bool
}

// AddRow adds a row to the buffered renderer.
func (r *Buffered) AddRow(table *Table, row Row) error {
	if r.rendered {
		return ErrAddAfterRender
	}
	// Store row for later rendering
	table.rows = append(table.rows, row)
	return nil
}

// Render renders the buffered content all at once.
func (r *Buffered) Render(table *Table) error {
	if r.rendered {
		return ErrTableAlreadyRendered
	}

	// Calculate column widths for auto-width columns
	table.CalculateColumnWidths()

	// Render all buffered content
	if err := table.RenderHeader(); err != nil {
		return err
	}

	for _, row := range table.rows {
		if err := table.RenderRow(row); err != nil {
			return err
		}
	}

	if err := table.RenderFooter(); err != nil {
		return err
	}

	r.rendered = true
	return nil
}

// IsRendered checks if the table has been rendered.
func (r *Buffered) IsRendered() bool {
	return r.rendered
}

// Streaming implements streaming rendering strategy.
type Streaming struct {
	rendered   bool
	headerDone bool
}

// AddRow adds a row to the streaming renderer.
func (r *Streaming) AddRow(table *Table, row Row) error {
	if r.rendered {
		return ErrAddAfterRender
	}

	if !r.headerDone {
		if err := table.RenderHeader(); err != nil {
			return err
		}
		r.headerDone = true
	}

	return table.RenderRow(row)
}

// Render renders the streaming content, typically just the footer.
func (r *Streaming) Render(table *Table) error {
	if r.rendered {
		return ErrTableAlreadyRendered
	}

	// For streaming mode, just render footer
	if err := table.RenderFooter(); err != nil {
		return err
	}

	r.rendered = true
	return nil
}

// IsRendered checks if the table has been rendered.
func (r *Streaming) IsRendered() bool {
	return r.rendered
}
