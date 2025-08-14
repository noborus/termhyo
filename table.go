package termhyo

import (
	"io"
	"strings"
)

// TableStyle represents the complete styling configuration for a table.
type TableStyle struct {
	BorderStyle BorderStyle // Border style configuration
	HeaderStyle HeaderStyle // Header styling configuration
}

// DefaultTableStyle returns the default table style.
func DefaultTableStyle() TableStyle {
	return TableStyle{
		BorderStyle: BoxDrawingStyle,
		HeaderStyle: HeaderStyle{},
	}
}

// WithBorderStyle creates a TableStyle with the specified border style.
func WithBorderStyle(borderStyle BorderStyle) TableStyle {
	return TableStyle{
		BorderStyle: borderStyle,
		HeaderStyle: HeaderStyle{},
	}
}

// WithHeaderStyle creates a TableStyle with the specified header style.
func WithHeaderStyle(headerStyle HeaderStyle) TableStyle {
	return TableStyle{
		BorderStyle: BoxDrawingStyle,
		HeaderStyle: headerStyle,
	}
}

// WithStyles creates a TableStyle with both border and header styles.
func WithStyles(borderStyle BorderStyle, headerStyle HeaderStyle) TableStyle {
	return TableStyle{
		BorderStyle: borderStyle,
		HeaderStyle: headerStyle,
	}
}

// Table represents the main table structure.
type Table struct {
	columns      []Column
	rows         []Row
	writer       io.Writer
	mode         RenderMode
	renderer     Renderer
	borderStyle  BorderStyle
	borderConfig BorderConfig
	align        bool // If false, skip alignment for all columns
	borders      map[string]string
	padding      int
	headerStyle  HeaderStyle // styling for header row
}

// GetBorderConfig returns the current border configuration.
func (t *Table) GetBorderConfig() BorderConfig {
	return t.borderConfig
}

// SetAlign sets whether to skip alignment for all columns.
func (t *Table) SetAlign(align bool) {
	t.align = align
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

// GetAlign returns the current align setting.
func (t *Table) GetAlign() bool {
	return t.align
}

// NewTable creates a new table with default styling.
func NewTable(writer io.Writer, columns []Column) *Table {
	return NewTableWithStyle(writer, columns, DefaultTableStyle())
}

// NewTableWithStyle creates a new table with specified styling.
func NewTableWithStyle(writer io.Writer, columns []Column, style TableStyle) *Table {
	borderConfig := getBorderConfig(style.BorderStyle)

	t := &Table{
		columns:      columns,
		writer:       writer,
		rows:         make([]Row, 0),
		padding:      1,
		align:        true, // Default to aligned columns
		borderStyle:  style.BorderStyle,
		borderConfig: borderConfig,
		borders:      borderConfig.Chars,
		headerStyle:  style.HeaderStyle,
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

// determineRenderMode decides whether to use buffered or streaming mode.
func (t *Table) determineRenderMode() RenderMode {
	hasAutoWidth := false

	for _, col := range t.columns {
		if col.Width == 0 {
			hasAutoWidth = true
			break
		}
	}

	// Use streaming mode if no auto-width calculation needed OR no alignment needed
	if !hasAutoWidth || !t.align {
		return StreamingMode
	}

	return BufferedMode
}

// AddRow adds a row to the table.
func (t *Table) AddRow(cells ...string) error {
	row := Row{
		Cells: make([]Cell, len(cells)),
	}

	for i, content := range cells {
		row.Cells[i] = Cell{Content: content}
	}

	return t.renderer.AddRow(t, row)
}

// AddRowCells adds a row with detailed cell configuration.
func (t *Table) AddRowCells(cells ...Cell) error {
	row := Row{Cells: cells}
	return t.renderer.AddRow(t, row)
}

// Render renders the complete table.
func (t *Table) Render() error {
	return t.renderer.Render(t)
}

// CalculateColumnWidths calculates optimal widths for auto-width columns.
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

		// Apply max width limit if set (before padding adjustment)
		if t.columns[colIndex].MaxWidth > 0 && maxWidth > t.columns[colIndex].MaxWidth {
			maxWidth = t.columns[colIndex].MaxWidth
		}

		t.columns[colIndex].Width = maxWidth
	}
}

// RenderHeader renders the table header.
func (t *Table) RenderHeader() error {
	// Top border (only if enabled)
	if t.borderConfig.Top {
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
			Align:   Center,
		}
	}

	// Use specialized header row rendering
	if err := t.RenderHeaderRow(headerRow); err != nil {
		return err
	}

	// Header separator (only if enabled)
	if t.borderConfig.Middle {
		return t.RenderBorderLine("middle")
	}

	return nil
}

// RenderHeaderRow renders a header row with full-line styling.
func (t *Table) RenderHeaderRow(row Row) error {
	var builder strings.Builder
	var stylePrefix, styleSuffix string

	// Apply header style to the entire line if configured
	if !t.headerStyle.isEmpty() {
		stylePrefix = t.headerStyle.getPrefix()
		styleSuffix = t.headerStyle.getSuffix()
	}

	// Start the line with style prefix
	builder.WriteString(stylePrefix)

	// Left border (only if enabled)
	if t.borderConfig.Left {
		builder.WriteString(t.borders["vertical"])
	}

	for i, col := range t.columns {
		var content string
		if i < len(row.Cells) {
			cell := row.Cells[i]
			content = cell.Content

			// Apply alignment if not disabled
			if t.align {
				align := col.Align
				if cell.Align != Default {
					align = cell.Align
				}
				content = t.formatCell(content, col.Width, align)
			}
		} else {
			if t.align {
				if !t.borderConfig.Padding {
					// No padding for empty cells
					content = strings.Repeat(" ", col.Width)
				} else {
					// Empty cell with padding using strings.Builder
					var cellBuilder strings.Builder
					paddingStr := strings.Repeat(" ", t.padding)
					cellBuilder.WriteString(paddingStr)
					cellBuilder.WriteString(strings.Repeat(" ", col.Width))
					cellBuilder.WriteString(paddingStr)
					content = cellBuilder.String()
				}
			}
		}

		builder.WriteString(content)

		// Add vertical separator between columns (only if enabled and not the last column)
		if t.borderConfig.Vertical && i < len(t.columns)-1 {
			builder.WriteString(t.borders["vertical"])
		}
	}

	// Right border (only if enabled)
	if t.borderConfig.Right {
		builder.WriteString(t.borders["vertical"])
	}

	// End the line with style suffix
	builder.WriteString(styleSuffix)
	builder.WriteString("\n")

	_, err := t.writer.Write([]byte(builder.String()))
	return err
}

// RenderRow renders a single row.
func (t *Table) RenderRow(row Row) error {
	var builder strings.Builder

	// Left border (only if enabled)
	if t.borderConfig.Left {
		builder.WriteString(t.borders["vertical"])
	}

	for i, col := range t.columns {
		var content string
		if i < len(row.Cells) {
			cell := row.Cells[i]
			content = cell.Content

			// Apply alignment if not disabled
			if t.align {
				align := col.Align
				if cell.Align != Default {
					align = cell.Align
				}
				content = t.formatCell(content, col.Width, align)
			}
		} else {
			if t.align {
				if !t.borderConfig.Padding {
					// No padding for empty cells
					content = strings.Repeat(" ", col.Width)
				} else {
					// Empty cell with padding using strings.Builder
					var cellBuilder strings.Builder
					paddingStr := strings.Repeat(" ", t.padding)
					cellBuilder.WriteString(paddingStr)
					cellBuilder.WriteString(strings.Repeat(" ", col.Width))
					cellBuilder.WriteString(paddingStr)
					content = cellBuilder.String()
				}
			}
		}

		builder.WriteString(content)

		// Add vertical separator between columns (only if enabled and not the last column)
		if t.borderConfig.Vertical && i < len(t.columns)-1 {
			builder.WriteString(t.borders["vertical"])
		}
	}

	// Right border (only if enabled)
	if t.borderConfig.Right {
		builder.WriteString(t.borders["vertical"])
	}

	builder.WriteString("\n")
	_, err := t.writer.Write([]byte(builder.String()))
	return err
}

// formatCell formats cell content with alignment and padding.
func (t *Table) formatCell(content string, width int, align Alignment) string {
	contentWidth := stringWidth(content)
	// Check if padding is disabled for this border style
	if !t.borderConfig.Padding {
		// No padding, use original behavior
		if contentWidth > width {
			content = truncateString(content, width)
		}
		return padString(content, width, align)
	}

	// For non-padding-disabled mode, width is the content width
	// Truncate if content is too long for the specified width
	if contentWidth > width {
		content = truncateString(content, width)
	}

	// Apply alignment to the content width
	paddedContent := padString(content, width, align)

	// Add padding spaces on both sides using strings.Builder for efficiency
	var builder strings.Builder
	paddingStr := strings.Repeat(" ", t.padding)
	builder.WriteString(paddingStr)
	builder.WriteString(paddedContent)
	builder.WriteString(paddingStr)
	return builder.String()
}

// RenderBorderLine renders horizontal border lines.
func (t *Table) RenderBorderLine(position string) error {
	var builder strings.Builder

	switch position {
	case "top":
		builder.WriteString(t.borders["top_left"])
	case "bottom":
		builder.WriteString(t.borders["bottom_left"])
	default: // middle
		builder.WriteString(t.borders["left_cross"])
	}

	for i, col := range t.columns {
		// Calculate the actual cell width (content + padding)
		cellWidth := col.Width
		if t.borderConfig.Padding {
			cellWidth += (t.padding * 2)
		}
		builder.WriteString(strings.Repeat(t.borders["horizontal"], cellWidth))

		if i < len(t.columns)-1 {
			switch position {
			case "top":
				builder.WriteString(t.borders["top_cross"])
			case "bottom":
				builder.WriteString(t.borders["bottom_cross"])
			default:
				builder.WriteString(t.borders["cross"])
			}
		}
	}

	switch position {
	case "top":
		builder.WriteString(t.borders["top_right"])
	case "bottom":
		builder.WriteString(t.borders["bottom_right"])
	default:
		builder.WriteString(t.borders["right_cross"])
	}

	builder.WriteString("\n")
	_, err := t.writer.Write([]byte(builder.String()))
	return err
}

// RenderFooter renders the table footer.
func (t *Table) RenderFooter() error {
	// Bottom border (only if enabled)
	if t.borderConfig.Bottom {
		return t.RenderBorderLine("bottom")
	}
	return nil
}

// SetRenderer allows setting a custom renderer.
func (t *Table) SetRenderer(renderer Renderer) {
	t.renderer = renderer
}

// SetBorderStyle changes the border style of the table.
func (t *Table) SetBorderStyle(style BorderStyle) {
	t.borderStyle = style
	t.borderConfig = getBorderConfig(style)
	t.borders = t.borderConfig.Chars
}

// GetBorderStyle returns the current border style.
func (t *Table) GetBorderStyle() BorderStyle {
	return t.borderStyle
}

// SetBorderConfig allows setting a custom border configuration.
func (t *Table) SetBorderConfig(config BorderConfig) {
	t.borderConfig = config
	t.borders = config.Chars
}

// SetHeaderStyle sets the styling for header row.
func (t *Table) SetHeaderStyle(style HeaderStyle) {
	t.headerStyle = style
}

// SetHeaderStyleWithoutSeparator sets the header style and disables the header separator line.
// This is a convenience method for the common use case of styled headers not needing separators.
func (t *Table) SetHeaderStyleWithoutSeparator(style HeaderStyle) {
	t.headerStyle = style
	// Disable header separator line since styled headers provide visual distinction
	t.borderConfig.Middle = false
	t.borders = t.borderConfig.Chars
}

// SetHeaderStyleWithoutBorders sets the header style and disables all horizontal borders.
// This creates a completely clean look with only the styled header for distinction.
func (t *Table) SetHeaderStyleWithoutBorders(style HeaderStyle) {
	t.headerStyle = style
	// Disable all horizontal borders for the cleanest look
	t.borderConfig.Top = false
	t.borderConfig.Middle = false
	t.borderConfig.Bottom = false
	t.borders = t.borderConfig.Chars
}

// SetHeaderStyleBorderless sets the header style and disables ALL borders including left/right.
// This creates the most minimal table with only styled header and column spacing.
func (t *Table) SetHeaderStyleBorderless(style HeaderStyle) {
	t.headerStyle = style
	// Disable all borders but keep internal vertical separators for column distinction
	t.borderConfig.Top = false
	t.borderConfig.Middle = false
	t.borderConfig.Bottom = false
	t.borderConfig.Left = false
	t.borderConfig.Right = false
	t.borderConfig.Vertical = true // Keep column separators
	t.borders = t.borderConfig.Chars
}

// SetHeaderStyleMinimal sets the header style and disables absolutely ALL borders and separators.
// This creates the most minimal possible table with only styled header and whitespace separation.
func (t *Table) SetHeaderStyleMinimal(style HeaderStyle) {
	t.headerStyle = style
	// Disable absolutely everything
	t.borderConfig.Top = false
	t.borderConfig.Middle = false
	t.borderConfig.Bottom = false
	t.borderConfig.Left = false
	t.borderConfig.Right = false
	t.borderConfig.Vertical = false // No column separators either
	t.borders = t.borderConfig.Chars
}

// GetHeaderStyle returns the current header style.
func (t *Table) GetHeaderStyle() HeaderStyle {
	return t.headerStyle
}
