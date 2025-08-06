package termhyo

import (
	"io"
	"strings"
)

// Table represents the main table structure
type Table struct {
	columns     []Column
	rows        []Row
	writer      io.Writer
	mode        RenderMode
	renderer    Renderer
	borderStyle BorderStyle

	// Style configuration
	borders map[string]string
	padding int
}

// NewTable creates a new table
func NewTable(writer io.Writer, columns []Column) *Table {
	return NewTableWithStyle(writer, columns, BoxDrawingStyle)
}

// NewTableWithStyle creates a new table with specified border style
func NewTableWithStyle(writer io.Writer, columns []Column, borderStyle BorderStyle) *Table {
	t := &Table{
		columns:     columns,
		writer:      writer,
		rows:        make([]Row, 0),
		padding:     1,
		borderStyle: borderStyle,
		borders:     getBorderChars(borderStyle),
	}

	// Determine render mode based on column configuration
	t.mode = t.determineRenderMode()

	// Set appropriate renderer based on mode
	if t.mode == StreamingMode {
		t.renderer = &Streaming{}
	} else {
		t.renderer = &Buffered{}
	}

	return t
}

// determineRenderMode decides whether to use buffered or streaming mode
func (t *Table) determineRenderMode() RenderMode {
	hasAutoWidth := false
	hasAlignment := false

	for _, col := range t.columns {
		if col.Width == 0 {
			hasAutoWidth = true
		}
		if !col.NoAlign {
			hasAlignment = true
		}
	}

	// Use streaming mode only if all widths are fixed and no alignment needed
	if !hasAutoWidth && !hasAlignment {
		return StreamingMode
	}

	return BufferedMode
}

// AddRow adds a row to the table
func (t *Table) AddRow(cells ...string) error {
	row := Row{
		Cells: make([]Cell, len(cells)),
	}

	for i, content := range cells {
		row.Cells[i] = Cell{Content: content}
	}

	return t.renderer.AddRow(t, row)
}

// AddRowCells adds a row with detailed cell configuration
func (t *Table) AddRowCells(cells ...Cell) error {
	row := Row{Cells: cells}
	return t.renderer.AddRow(t, row)
}

// Render renders the complete table
func (t *Table) Render() error {
	return t.renderer.Render(t)
}

// CalculateColumnWidths calculates optimal widths for auto-width columns
func (t *Table) CalculateColumnWidths() {
	for i, col := range t.columns {
		if col.Width == 0 { // auto-width column
			maxWidth := len(col.Title) // start with header width

			// Check all data rows
			for _, row := range t.rows {
				if i < len(row.Cells) {
					contentWidth := len(row.Cells[i].Content)
					if contentWidth > maxWidth {
						maxWidth = contentWidth
					}
				}
			}

			// Apply max width limit if set
			if col.MaxWidth > 0 && maxWidth > col.MaxWidth {
				maxWidth = col.MaxWidth
			}

			t.columns[i].Width = maxWidth
		}
	}
}

// RenderHeader renders the table header
func (t *Table) RenderHeader() error {
	// Top border
	if err := t.RenderBorderLine("top"); err != nil {
		return err
	}

	// Header row
	headerRow := Row{
		Cells: make([]Cell, len(t.columns)),
	}
	for i, col := range t.columns {
		headerRow.Cells[i] = Cell{
			Content: col.Title,
			Align:   "center",
		}
	}

	if err := t.RenderRow(headerRow); err != nil {
		return err
	}

	// Header separator
	return t.RenderBorderLine("middle")
}

// RenderRow renders a single row
func (t *Table) RenderRow(row Row) error {
	line := t.borders["vertical"]

	for i, col := range t.columns {
		var content string
		if i < len(row.Cells) {
			cell := row.Cells[i]
			content = cell.Content

			// Apply alignment if not disabled
			if !col.NoAlign {
				align := col.Align
				if cell.Align != "" {
					align = cell.Align
				}
				content = t.formatCell(content, col.Width, align)
			}
		} else {
			if !col.NoAlign {
				content = strings.Repeat(" ", col.Width)
			}
		}

		line += content + t.borders["vertical"]
	}

	line += "\n"
	_, err := t.writer.Write([]byte(line))
	return err
}

// formatCell formats cell content with alignment and padding
func (t *Table) formatCell(content string, width int, align string) string {
	// Truncate if too long
	if len(content) > width {
		if width > 3 {
			content = content[:width-3] + "..."
		} else {
			content = content[:width]
		}
	}

	padding := width - len(content)

	switch align {
	case "right":
		return strings.Repeat(" ", padding) + content
	case "center":
		leftPad := padding / 2
		rightPad := padding - leftPad
		return strings.Repeat(" ", leftPad) + content + strings.Repeat(" ", rightPad)
	default: // left
		return content + strings.Repeat(" ", padding)
	}
}

// RenderBorderLine renders horizontal border lines
func (t *Table) RenderBorderLine(position string) error {
	var line string

	switch position {
	case "top":
		line = t.borders["top_left"]
	case "bottom":
		line = t.borders["bottom_left"]
	default: // middle
		line = t.borders["left_cross"]
	}

	for i, col := range t.columns {
		line += strings.Repeat(t.borders["horizontal"], col.Width)

		if i < len(t.columns)-1 {
			switch position {
			case "top":
				line += t.borders["top_cross"]
			case "bottom":
				line += t.borders["bottom_cross"]
			default:
				line += t.borders["cross"]
			}
		}
	}

	switch position {
	case "top":
		line += t.borders["top_right"]
	case "bottom":
		line += t.borders["bottom_right"]
	default:
		line += t.borders["right_cross"]
	}

	line += "\n"
	_, err := t.writer.Write([]byte(line))
	return err
}

// RenderFooter renders the table footer
func (t *Table) RenderFooter() error {
	return t.RenderBorderLine("bottom")
}

// SetRenderer allows setting a custom renderer
func (t *Table) SetRenderer(renderer Renderer) {
	t.renderer = renderer
}

// SetBorderStyle changes the border style of the table
func (t *Table) SetBorderStyle(style BorderStyle) {
	t.borderStyle = style
	t.borders = getBorderChars(style)
}

// GetBorderStyle returns the current border style
func (t *Table) GetBorderStyle() BorderStyle {
	return t.borderStyle
}
