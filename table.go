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
	noAlign      bool // skip alignment entirely for all columns

	// Style configuration
	borders map[string]string
	padding int
}

// GetBorderConfig returns the current border configuration
func (t *Table) GetBorderConfig() BorderConfig {
	return t.borderConfig
}

// SetNoAlign sets whether to skip alignment for all columns
func (t *Table) SetNoAlign(noAlign bool) {
	t.noAlign = noAlign
	// Recalculate render mode when alignment setting changes
	t.mode = t.determineRenderMode()

	// Update renderer based on new mode
	if t.borderStyle == MarkdownStyle {
		t.renderer = &MarkdownRenderer{}
	} else if t.mode == StreamingMode {
		t.renderer = &Streaming{}
	} else {
		t.renderer = &Buffered{}
	}
}

// GetNoAlign returns the current no-align setting
func (t *Table) GetNoAlign() bool {
	return t.noAlign
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

	for _, col := range t.columns {
		if col.Width == 0 {
			hasAutoWidth = true
			break
		}
	}

	// Use streaming mode if no auto-width calculation needed OR no alignment needed
	if !hasAutoWidth || t.noAlign {
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
	// Early return if no auto-width columns
	autoWidthColumns := make([]int, 0, len(t.columns)) // Track auto-width column indices
	for i, col := range t.columns {
		if col.Width == 0 {
			autoWidthColumns = append(autoWidthColumns, i)
		}
	}
	if len(autoWidthColumns) == 0 {
		return // All columns have fixed widths, no calculation needed
	}

	// Initialize max widths with header widths for auto-width columns
	maxWidths := make([]int, len(t.columns))
	for _, colIndex := range autoWidthColumns {
		maxWidths[colIndex] = stringWidth(t.columns[colIndex].Title)
	}

	// Check all data rows for accurate width calculation (row-oriented for better cache efficiency)
	for _, row := range t.rows {
		for _, colIndex := range autoWidthColumns { // Only process auto-width columns
			if colIndex < len(row.Cells) {
				contentWidth := stringWidth(row.Cells[colIndex].Content)
				if contentWidth > maxWidths[colIndex] {
					maxWidths[colIndex] = contentWidth
				}
			}
		}
	}

	// Apply final width calculations
	for _, colIndex := range autoWidthColumns {
		maxWidth := maxWidths[colIndex]

		// Add padding to the calculated width (padding on both sides) only if padding is enabled
		if !t.borderConfig.DisablePadding {
			maxWidth += (t.padding * 2)
		}

		// Apply max width limit if set (after padding adjustment)
		if t.columns[colIndex].MaxWidth > 0 && maxWidth > t.columns[colIndex].MaxWidth {
			maxWidth = t.columns[colIndex].MaxWidth
		}

		t.columns[colIndex].Width = maxWidth
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
			if !t.noAlign {
				align := col.Align
				if cell.Align != "" {
					align = cell.Align
				}
				content = t.formatCell(content, col.Width, align)
			}
		} else {
			if !t.noAlign {
				if t.borderConfig.DisablePadding {
					// No padding for empty cells
					content = strings.Repeat(" ", col.Width)
				} else {
					// Empty cell with padding
					paddingStr := strings.Repeat(" ", t.padding)
					effectiveWidth := col.Width - (t.padding * 2)
					if effectiveWidth < 0 {
						effectiveWidth = 0
					}
					content = paddingStr + strings.Repeat(" ", effectiveWidth) + paddingStr
				}
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
	// Check if padding is disabled for this border style
	if t.borderConfig.DisablePadding {
		// No padding, use original behavior
		if stringWidth(content) > width {
			content = truncateString(content, width)
		}
		return padString(content, width, align)
	}

	// Calculate effective width (subtract padding from both sides)
	effectiveWidth := width - (t.padding * 2)
	if effectiveWidth < 0 {
		effectiveWidth = 0
	}

	// Truncate if too long for effective width
	if stringWidth(content) > effectiveWidth {
		content = truncateString(content, effectiveWidth)
	}

	// Apply alignment to effective width
	paddedContent := padString(content, effectiveWidth, align)

	// Add padding spaces on both sides
	paddingStr := strings.Repeat(" ", t.padding)
	return paddingStr + paddedContent + paddingStr
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
