package termhyo

import "fmt"

// Renderer defines the interface for different rendering strategies
type Renderer interface {
	AddRow(table *Table, row Row) error
	Render(table *Table) error
	IsRendered() bool
}

// RenderMode defines when to start rendering
type RenderMode int

const (
	// BufferedMode: collect all rows before rendering (for auto-width)
	BufferedMode RenderMode = iota
	// StreamingMode: render immediately (for fixed-width)
	StreamingMode
)

// Buffered implements buffered rendering strategy
type Buffered struct {
	rendered bool
}

func (r *Buffered) AddRow(table *Table, row Row) error {
	if r.rendered {
		return fmt.Errorf("cannot add rows after table is rendered")
	}
	// Store row for later rendering
	table.rows = append(table.rows, row)
	return nil
}

func (r *Buffered) Render(table *Table) error {
	if r.rendered {
		return fmt.Errorf("table already rendered")
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

func (r *Buffered) IsRendered() bool {
	return r.rendered
}

// Streaming implements streaming rendering strategy
type Streaming struct {
	rendered   bool
	headerDone bool
}

func (r *Streaming) AddRow(table *Table, row Row) error {
	if r.rendered {
		return fmt.Errorf("cannot add rows after table is rendered")
	}

	if !r.headerDone {
		if err := table.RenderHeader(); err != nil {
			return err
		}
		r.headerDone = true
	}

	return table.RenderRow(row)
}

func (r *Streaming) Render(table *Table) error {
	if r.rendered {
		return fmt.Errorf("table already rendered")
	}

	// For streaming mode, just render footer
	if err := table.RenderFooter(); err != nil {
		return err
	}

	r.rendered = true
	return nil
}

func (r *Streaming) IsRendered() bool {
	return r.rendered
}
