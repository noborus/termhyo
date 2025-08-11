# Header Styles Example

This example demonstrates advanced header styling capabilities with ANSI escape sequences.

## Running the Example

```bash
go run main.go
```

## What This Example Shows

- **ANSI Color Support**: Using foreground and background colors
- **Text Formatting**: Bold, italic, underline effects
- **Border Control**: Progressive border removal for cleaner looks (using `Top`, `Middle`, `Bottom` fields)
- **Convenience Methods**: Using helper methods for common patterns
- **True Color Support**: 24-bit color examples

## Features Demonstrated

1. **Basic Header Styling**: Bold and colored headers
2. **Background Colors**: Headers with background colors
3. **Border Removal**: Removing separator lines and borders (e.g. `Middle: false`, `Top: false`, `Bottom: false`)
4. **Minimal Tables**: Clean tables with only header styling (e.g. `Top: false`, `Bottom: false`, `Left: false`, `Right: false`, `Vertical: false`)
5. **Color Variations**: Different color combinations
6. **True Color**: RGB color specification

## Convenience Methods & BorderConfig Usage

The example shows these convenience methods:

- `SetHeaderStyleWithoutSeparator()` - Removes the header separator line
- `SetHeaderStyleWithoutBorders()` - Removes all horizontal borders
- `SetHeaderStyleBorderless()` - Removes all borders except column separators
- `SetHeaderStyleMinimal()` - Creates the most minimal table

You can also directly control borders using the new `BorderConfig` fields:

```go
borderConfig := table.GetBorderConfig()
borderConfig.Top = false      // Hide top border
borderConfig.Middle = false   // Hide header separator
borderConfig.Bottom = false   // Hide bottom border
borderConfig.Left = false     // Hide left border
borderConfig.Right = false    // Hide right border
borderConfig.Vertical = false // Hide all internal vertical separators
table.SetBorderConfig(borderConfig)
```

This demonstrates how header styling can replace traditional table borders for a modern look.
