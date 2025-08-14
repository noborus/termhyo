package termhyo

import (
	"fmt"
	"strings"
)

// MarkdownRenderer implements Markdown table format with streaming support.
type MarkdownRenderer struct {
	rendered     bool
	bufferedRows []Row // Buffer rows for width calculation
}

// hasAutoWidth checks if any columns have auto width.
func hasAutoWidth(table *Table) bool {
	if !table.align {
		return false // No auto width in streaming mode
	}
	for _, col := range table.columns {
		if col.Width == 0 {
			return true
		}
	}
	return false
}

// AddRow adds a row for markdown rendering (buffered mode for width calculation).
func (r *MarkdownRenderer) AddRow(_ *Table, row Row) error {
	if r.rendered {
		return fmt.Errorf("cannot add rows after table is rendered")
	}

	// Buffer the row for width calculation
	r.bufferedRows = append(r.bufferedRows, row)

	// Don't render immediately - wait for Render() call
	return nil
}

// Render renders any remaining content (for streaming, this is just cleanup).
func (r *MarkdownRenderer) Render(table *Table) error {
	if r.rendered {
		return fmt.Errorf("table already rendered")
	}

	// Calculate column widths if needed using all buffered rows
	if hasAutoWidth(table) {
		// Temporarily copy buffered rows to table for width calculation
		originalRows := table.rows
		table.rows = r.bufferedRows
		table.CalculateColumnWidths()
		table.rows = originalRows // Restore original rows
	}

	// Render header and separator
	if err := r.renderMarkdownHeader(table); err != nil {
		return err
	}
	if err := r.renderMarkdownSeparator(table); err != nil {
		return err
	}

	// Render all buffered rows
	for _, row := range r.bufferedRows {
		if err := r.renderMarkdownRow(table, row); err != nil {
			return err
		}
	}

	r.rendered = true
	return nil
}

// IsRendered returns whether the table has been rendered.
func (r *MarkdownRenderer) IsRendered() bool {
	return r.rendered
}

// renderMarkdownHeader renders the header row in Markdown format.
func (r *MarkdownRenderer) renderMarkdownHeader(table *Table) error {
	var line string
	var stylePrefix, styleSuffix string

	// Apply header style to the entire line if configured
	if !table.headerStyle.isEmpty() {
		stylePrefix = table.headerStyle.getPrefix()
		styleSuffix = table.headerStyle.getSuffix()
	}

	// Start the line with style prefix
	line = stylePrefix + "|"

	for _, col := range table.columns {
		// Apply alignment to header content (headers are typically centered)
		content := col.Title
		if table.align {
			content = table.formatCell(col.Title, col.Width, Center)
		}
		line += content + "|"
	}

	// End the line with style suffix
	line += styleSuffix + "\n"
	_, err := table.writer.Write([]byte(line))
	return err
}

// renderMarkdownSeparator renders the separator row with alignment indicators.
func (r *MarkdownRenderer) renderMarkdownSeparator(table *Table) error {
	line := "|"

	for _, col := range table.columns {
		separatorWidth := max(col.Width, 1)
		if table.borderConfig.Padding {
			separatorWidth += (table.padding * 2)
		}
		separator := r.getAlignmentSeparator(col.Align, separatorWidth)
		line += separator + "|"
	}

	line += "\n"
	_, err := table.writer.Write([]byte(line))
	return err
}

// renderMarkdownRow renders a data row in Markdown format.
func (r *MarkdownRenderer) renderMarkdownRow(table *Table, row Row) error {
	line := "|"

	// Ensure row.Cells has at least as many elements as table.columns
	cells := row.Cells
	if len(cells) < len(table.columns) {
		// Pad with empty cells if necessary
		for i := len(cells); i < len(table.columns); i++ {
			cells = append(cells, Cell{Content: ""})
		}
	}

	for i, col := range table.columns {
		if !table.align {
			line += cells[i].Content + "|"
			continue // Skip alignment if noAlign is set
		}

		var content string
		// Apply column alignment to cell content
		cellAlign := col.Align
		if cells[i].Align != Default {
			cellAlign = cells[i].Align // Cell-specific alignment overrides column alignment
		}
		// Format cell content with alignment
		content = table.formatCell(cells[i].Content, col.Width, cellAlign)
		line += content + "|"
	}

	line += "\n"
	_, err := table.writer.Write([]byte(line))
	return err
}

// getAlignmentSeparator returns the separator string with alignment indicators.
func (r *MarkdownRenderer) getAlignmentSeparator(align Alignment, width int) string {
	switch align {
	case Right:
		if width <= 1 {
			return ":"
		}
		return strings.Repeat("-", width-1) + ":"
	case Center:
		if width <= 2 {
			return "::"
		}
		return ":" + strings.Repeat("-", width-2) + ":"
	default: // Left, Default, or no alignment
		return strings.Repeat("-", width)
	}
}
