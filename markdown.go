package termhyo

import (
	"fmt"
	"strings"
)

// MarkdownRenderer implements Markdown table format with streaming support
type MarkdownRenderer struct {
	rendered   bool
	headerDone bool
}

// hasAutoWidth checks if any columns have auto width
func hasAutoWidth(table *Table) bool {
	for _, col := range table.columns {
		if col.Width == 0 {
			return true
		}
	}
	return false
}

// AddRow adds a row for markdown rendering (streaming mode)
func (r *MarkdownRenderer) AddRow(table *Table, row Row) error {
	if r.rendered {
		return fmt.Errorf("cannot add rows after table is rendered")
	}

	// Render header and separator on first row
	if !r.headerDone {
		// Calculate column widths if needed
		if hasAutoWidth(table) {
			table.rows = append(table.rows, row) // Temporarily add for width calculation
			table.CalculateColumnWidths()
			table.rows = table.rows[:len(table.rows)-1] // Remove temporary row
		}

		if err := r.renderMarkdownHeader(table); err != nil {
			return err
		}
		if err := r.renderMarkdownSeparator(table); err != nil {
			return err
		}
		r.headerDone = true
	}

	// Render the data row immediately
	return r.renderMarkdownRow(table, row)
}

// Render renders any remaining content (for streaming, this is just cleanup)
func (r *MarkdownRenderer) Render(table *Table) error {
	if r.rendered {
		return fmt.Errorf("table already rendered")
	}

	// For streaming mode, if no rows were added, render header and separator
	if !r.headerDone {
		// Calculate column widths if needed
		if hasAutoWidth(table) {
			table.CalculateColumnWidths()
		}

		if err := r.renderMarkdownHeader(table); err != nil {
			return err
		}
		if err := r.renderMarkdownSeparator(table); err != nil {
			return err
		}
	}

	r.rendered = true
	return nil
}

// IsRendered returns whether the table has been rendered
func (r *MarkdownRenderer) IsRendered() bool {
	return r.rendered
}

// renderMarkdownHeader renders the header row in Markdown format
func (r *MarkdownRenderer) renderMarkdownHeader(table *Table) error {
	line := "|"

	for _, col := range table.columns {
		// Apply alignment to header content (headers are typically centered)
		content := table.formatCell(col.Title, col.Width, "center")
		line += content + "|"
	}

	line += "\n"
	_, err := table.writer.Write([]byte(line))
	return err
}

// renderMarkdownSeparator renders the separator row with alignment indicators
func (r *MarkdownRenderer) renderMarkdownSeparator(table *Table) error {
	line := "|"

	for _, col := range table.columns {
		separator := r.getAlignmentSeparator(col.Align, col.Width)
		line += separator + "|"
	}

	line += "\n"
	_, err := table.writer.Write([]byte(line))
	return err
}

// renderMarkdownRow renders a data row in Markdown format
func (r *MarkdownRenderer) renderMarkdownRow(table *Table, row Row) error {
	line := "|"

	for i, col := range table.columns {
		var content string
		if i < len(row.Cells) {
			// Apply column alignment to cell content
			cellAlign := col.Align
			if row.Cells[i].Align != "" {
				cellAlign = row.Cells[i].Align // Cell-specific alignment overrides column alignment
			}
			content = table.formatCell(row.Cells[i].Content, col.Width, cellAlign)
		} else {
			content = table.formatCell("", col.Width, col.Align)
		}
		line += content + "|"
	}

	line += "\n"
	_, err := table.writer.Write([]byte(line))
	return err
}

// getAlignmentSeparator returns the separator string with alignment indicators
func (r *MarkdownRenderer) getAlignmentSeparator(align string, width int) string {
	switch align {
	case "right":
		if width <= 1 {
			return ":"
		}
		return strings.Repeat("-", width-1) + ":"
	case "center":
		if width <= 2 {
			return "::"
		}
		return ":" + strings.Repeat("-", width-2) + ":"
	default: // left or no alignment
		return strings.Repeat("-", width)
	}
}
