package termhyo

import (
	"errors"
	"io"
	"strings"
)

var (
	// ErrNoColumns is returned when trying to render a table with no columns defined.
	ErrNoColumns = errors.New("no columns defined")
	// ErrNoHeader is returned when trying to render a table without a header.
	ErrNoHeader = errors.New("header is required")
	// ErrNoMarkdownHeader is returned when trying to render a markdown table without a header.
	ErrNoMarkdownHeader = errors.New("markdown table requires at least one non-empty header")
	// ErrTableAlreadyRendered is returned when trying to render a table that has already been rendered.
	ErrTableAlreadyRendered = errors.New("table has already been rendered")
	// ErrAddAfterRender is returned when trying to add a row after rendering.
	ErrAddAfterRender = errors.New("cannot add row after table has been rendered")
)

// TableOption is a functional option for configuring Table.
type TableOption func(*Table)

// Table represents the main table structure.
type Table struct {
	columns      []Column
	rows         []Row
	writer       io.Writer
	mode         RenderMode
	renderer     Renderer
	borderStyle  BorderStyle
	borderConfig TableBorderConfig
	autoAlign    bool // If false, skip alignment for all columns
	borders      map[string]string
	padding      int
	headerStyle  HeaderStyle // styling for header row
}

// NewTable creates a new table with the given columns and optional configuration.
//
// This function uses the Functional Option Pattern:
//
//	table := termhyo.NewTable(os.Stdout, columns, termhyo.Border(termhyo.BoxDrawingStyle), termhyo.Header(headerStyle), ...)
//
// You can specify border style, header style, alignment, etc. by passing option functions.
//
// Example:
//
//	table := termhyo.NewTable(os.Stdout, columns, termhyo.Border(termhyo.DoubleStyle), termhyo.AutoAlign(false))
//
// This is the recommended way to create and configure tables in termhyo.
func NewTable(writer io.Writer, columns []Column, opts ...TableOption) *Table {
	borderConfig := GetBorderConfig(BoxDrawingStyle)

	t := &Table{
		columns:      columns,
		writer:       writer,
		rows:         make([]Row, 0),
		padding:      1,
		autoAlign:    true, // Default to auto-aligning columns
		borderStyle:  BoxDrawingStyle,
		borderConfig: borderConfig,
		borders:      borderConfig.Chars,
		headerStyle:  HeaderStyle{},
	}

	// Apply options
	for _, opt := range opts {
		opt(t)
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
	if !hasAutoWidth || !t.autoAlign {
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

// RenderHeader renders the table header row, including the top border and header separator line if enabled.
func (t *Table) RenderHeader() error {
	if len(t.columns) == 0 {
		return ErrNoColumns
	}
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

	// Cache vertical border string
	vertical := t.borders["vertical"]

	// Left border (only if enabled)
	if t.borderConfig.Left {
		builder.WriteString(vertical)
	}

	for i, col := range t.columns {
		content := t.getCellContent(row, i, col)
		builder.WriteString(content)

		// Add vertical separator between columns (only if enabled and not the last column)
		if t.borderConfig.Vertical && i < len(t.columns)-1 {
			builder.WriteString(vertical)
		}
	}

	// Right border (only if enabled)
	if t.borderConfig.Right {
		builder.WriteString(vertical)
	}

	// End the line with style suffix
	builder.WriteString(styleSuffix)
	builder.WriteString("\n")

	_, err := t.writer.Write([]byte(builder.String()))
	return err
}

// getCellContent returns the formatted content for a cell in a header row and column.
func (t *Table) getCellContent(row Row, i int, col Column) string {
	if i < len(row.Cells) {
		cell := row.Cells[i]
		if !t.autoAlign {
			return cell.Content // No alignment, return raw content
		}

		align := col.Align
		if cell.Align != Default {
			align = cell.Align
		}
		return t.formatCell(cell.Content, col.Width, align)
	}

	if !t.autoAlign {
		return ""
	}
	if !t.borderConfig.Padding {
		// No padding for empty cells
		return strings.Repeat(" ", col.Width)
	}
	// Empty cell with padding using strings.Builder
	var builder strings.Builder
	paddingStr := strings.Repeat(" ", t.padding)
	builder.WriteString(paddingStr)
	builder.WriteString(strings.Repeat(" ", col.Width))
	builder.WriteString(paddingStr)
	return builder.String()
}

// RenderRow renders a single row.
func (t *Table) RenderRow(row Row) error {
	var builder strings.Builder

	// Cache vertical border string
	vertical := t.borders["vertical"]

	// Left border (only if enabled)
	if t.borderConfig.Left {
		builder.WriteString(vertical)
	}

	for i, col := range t.columns {
		content := t.getCellContent(row, i, col)
		builder.WriteString(content)

		// Add vertical separator between columns (only if enabled and not the last column)
		if t.borderConfig.Vertical && i < len(t.columns)-1 {
			builder.WriteString(vertical)
		}
	}

	// Right border (only if enabled)
	if t.borderConfig.Right {
		builder.WriteString(vertical)
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

	// When padding is enabled, the cell width is "content width + padding on both sides"
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

	// left border (only if enabled)
	if t.borderConfig.Left {
		switch position {
		case "top":
			builder.WriteString(t.borders["top_left"])
		case "bottom":
			builder.WriteString(t.borders["bottom_left"])
		default:
			builder.WriteString(t.borders["left_cross"])
		}
	}

	for i, col := range t.columns {
		// Calculate the actual cell width (content + padding)
		cellWidth := col.Width
		if t.borderConfig.Padding {
			cellWidth += (t.padding * 2)
		}
		builder.WriteString(strings.Repeat(t.borders["horizontal"], cellWidth))

		// Draw vertical separator between columns only if enabled
		if t.borderConfig.Vertical && i < len(t.columns)-1 {
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

	// right border (only if enabled)
	if t.borderConfig.Right {
		switch position {
		case "top":
			builder.WriteString(t.borders["top_right"])
		case "bottom":
			builder.WriteString(t.borders["bottom_right"])
		default:
			builder.WriteString(t.borders["right_cross"])
		}
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

// Border sets the border style (option).
func Border(style BorderStyle) TableOption {
	return func(t *Table) {
		t.borderStyle = style
		t.borderConfig = GetBorderConfig(style)
		t.borders = t.borderConfig.Chars
	}
}

// BorderConfig sets a custom border configuration (option).
//
// Example:
//
//	cfg := termhyo.GetBorderConfig(termhyo.BoxDrawingStyle)
//	cfg.Left = false
//	table := termhyo.NewTable(os.Stdout, columns, termhyo.BorderConfig(cfg))
func BorderConfig(cfg TableBorderConfig) TableOption {
	return func(t *Table) {
		t.borderConfig = cfg
		t.borders = cfg.Chars
	}
}

// Header sets the header style (option).
func Header(style HeaderStyle) TableOption {
	return func(t *Table) {
		t.headerStyle = style
	}
}

// AutoAlign sets the align flag (option).
func AutoAlign(autoAlign bool) TableOption {
	return func(t *Table) {
		t.autoAlign = autoAlign
	}
}

// GetBorderConfig returns the current border configuration.
func (t *Table) GetBorderConfig() TableBorderConfig {
	return t.borderConfig
}

// SetAutoAlign sets whether to skip alignment for all columns.
func (t *Table) SetAutoAlign(autoAlign bool) {
	t.autoAlign = autoAlign
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

// GetAutoAlign returns the current auto-align setting.
func (t *Table) GetAutoAlign() bool {
	return t.autoAlign
}

// SetRenderer allows setting a custom renderer.
func (t *Table) SetRenderer(renderer Renderer) {
	t.renderer = renderer
}

// SetBorderStyle changes the border style of the table.
func (t *Table) SetBorderStyle(style BorderStyle) {
	t.borderStyle = style
	t.borderConfig = GetBorderConfig(style)
	t.borders = t.borderConfig.Chars
}

// GetBorderStyle returns the current border style.
func (t *Table) GetBorderStyle() BorderStyle {
	return t.borderStyle
}

// SetBorderConfig allows setting a custom border configuration.
func (t *Table) SetBorderConfig(config TableBorderConfig) {
	t.borderConfig = config
	t.borders = config.Chars
}

// SetHeaderStyle sets the styling for header row.
func (t *Table) SetHeaderStyle(style HeaderStyle) {
	t.headerStyle = style
}

// GetHeaderStyle returns the current header style.
func (t *Table) GetHeaderStyle() HeaderStyle {
	return t.headerStyle
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
