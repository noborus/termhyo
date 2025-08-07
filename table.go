package termhyo

import (
	"io"
	"strings"
)

// Table represents the main table structure
type Table struct {
	columns      []Column
	rows         []Row
	writer       io.Writer
	mode         RenderMode
	renderer     Renderer
	borderStyle  BorderStyle
	borderConfig BorderConfig

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
	borderConfig := getBorderConfig(borderStyle)

	t := &Table{
		columns:      columns,
		writer:       writer,
		rows:         make([]Row, 0),
		padding:      1,
		borderStyle:  borderStyle,
		borderConfig: borderConfig,
		borders:      borderConfig.Chars,
	}

	// Determine render mode based on column configuration
	t.mode = t.determineRenderMode()

	// Set appropriate renderer based on mode and style
	if t.borderStyle == MarkdownStyle {
		t.renderer = &MarkdownRenderer{}
	} else if t.mode == StreamingMode {
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
			maxWidth := stringWidth(col.Title) // start with header width

			// Check all data rows
			for _, row := range t.rows {
				if i < len(row.Cells) {
					contentWidth := stringWidth(row.Cells[i].Content)
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
	// Top border (only if enabled)
	if !t.borderConfig.DisableTop {
		if err := t.RenderBorderLine("top"); err != nil {
			return err
		}
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

	// Header separator (only if enabled)
	if !t.borderConfig.DisableMiddle {
		return t.RenderBorderLine("middle")
	}

	return nil
}

// RenderRow renders a single row
func (t *Table) RenderRow(row Row) error {
	var line string

	// Left border (only if enabled)
	if !t.borderConfig.DisableLeft {
		line = t.borders["vertical"]
	}

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

		line += content

		// Add vertical separator between columns (only if enabled and not the last column)
		if !t.borderConfig.DisableVertical && i < len(t.columns)-1 {
			line += t.borders["vertical"]
		}
	}

	// Right border (only if enabled)
	if !t.borderConfig.DisableRight {
		line += t.borders["vertical"]
	}

	line += "\n"
	_, err := t.writer.Write([]byte(line))
	return err
}

// formatCell formats cell content with alignment and padding
func (t *Table) formatCell(content string, width int, align string) string {
	// Truncate if too long
	if stringWidth(content) > width {
		content = truncateString(content, width)
	}

	return padString(content, width, align)
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
	// Bottom border (only if enabled)
	if !t.borderConfig.DisableBottom {
		return t.RenderBorderLine("bottom")
	}
	return nil
}

// SetRenderer allows setting a custom renderer
func (t *Table) SetRenderer(renderer Renderer) {
	t.renderer = renderer
}

// SetBorderStyle changes the border style of the table
func (t *Table) SetBorderStyle(style BorderStyle) {
	t.borderStyle = style
	t.borderConfig = getBorderConfig(style)
	t.borders = t.borderConfig.Chars
}

// GetBorderStyle returns the current border style
func (t *Table) GetBorderStyle() BorderStyle {
	return t.borderStyle
}

// SetBorderConfig allows setting a custom border configuration
func (t *Table) SetBorderConfig(config BorderConfig) {
	t.borderConfig = config
	t.borders = config.Chars
}

// GetBorderConfig returns the current border configuration
func (t *Table) GetBorderConfig() BorderConfig {
	return t.borderConfig
}
