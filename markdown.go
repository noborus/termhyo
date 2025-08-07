package termhyo

import (
	"fmt"
	"strings"
)

// MarkdownRenderer implements Markdown table format
type MarkdownRenderer struct {
	table    *Table
	rendered bool
}

// AddRow adds a row for markdown rendering
func (r *MarkdownRenderer) AddRow(table *Table, row Row) error {
	if r.rendered {
		return fmt.Errorf("cannot add rows after table is rendered")
	}
	r.table = table
	table.rows = append(table.rows, row)
	return nil
}

// Render renders the table in Markdown format
func (r *MarkdownRenderer) Render(table *Table) error {
	if r.rendered {
		return fmt.Errorf("table already rendered")
	}

	r.table = table

	// Calculate column widths for proper alignment
	table.CalculateColumnWidths()

	// Render header row
	if err := r.renderMarkdownHeader(); err != nil {
		return err
	}

	// Render separator row with alignment indicators
	if err := r.renderMarkdownSeparator(); err != nil {
		return err
	}

	// Render data rows
	for _, row := range table.rows {
		if err := r.renderMarkdownRow(row); err != nil {
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
func (r *MarkdownRenderer) renderMarkdownHeader() error {
	line := "|"

	for _, col := range r.table.columns {
		// Apply alignment to header content (headers are typically centered)
		content := r.table.formatCell(col.Title, col.Width, "center")
		line += content + "|"
	}

	line += "\n"
	_, err := r.table.writer.Write([]byte(line))
	return err
}

// renderMarkdownSeparator renders the separator row with alignment indicators
func (r *MarkdownRenderer) renderMarkdownSeparator() error {
	line := "|"

	for _, col := range r.table.columns {
		separator := r.getAlignmentSeparator(col.Align, col.Width)
		line += separator + "|"
	}

	line += "\n"
	_, err := r.table.writer.Write([]byte(line))
	return err
}

// renderMarkdownRow renders a data row in Markdown format
func (r *MarkdownRenderer) renderMarkdownRow(row Row) error {
	line := "|"

	for i, col := range r.table.columns {
		var content string
		if i < len(row.Cells) {
			// Apply column alignment to cell content
			cellAlign := col.Align
			if row.Cells[i].Align != "" {
				cellAlign = row.Cells[i].Align // Cell-specific alignment overrides column alignment
			}
			content = r.table.formatCell(row.Cells[i].Content, col.Width, cellAlign)
		} else {
			content = r.table.formatCell("", col.Width, col.Align)
		}
		line += content + "|"
	}

	line += "\n"
	_, err := r.table.writer.Write([]byte(line))
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
