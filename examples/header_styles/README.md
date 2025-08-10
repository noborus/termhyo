# Header Styles Example

This example demonstrates advanced header styling capabilities with ANSI escape sequences.

## Running the Example

```bash
go run main.go
```

## What This Example Shows

- **ANSI Color Support**: Using foreground and background colors
- **Text Formatting**: Bold, italic, underline effects
- **Border Control**: Progressive border removal for cleaner looks
- **Convenience Methods**: Using helper methods for common patterns
- **True Color Support**: 24-bit color examples

## Features Demonstrated

1. **Basic Header Styling**: Bold and colored headers
2. **Background Colors**: Headers with background colors
3. **Border Removal**: Removing separator lines and borders
4. **Minimal Tables**: Clean tables with only header styling
5. **Color Variations**: Different color combinations
6. **True Color**: RGB color specification

## Convenience Methods

The example shows these convenience methods:

- `SetHeaderStyleWithoutSeparator()` - Removes the header separator line
- `SetHeaderStyleWithoutBorders()` - Removes all horizontal borders
- `SetHeaderStyleBorderless()` - Removes all borders except column separators
- `SetHeaderStyleMinimal()` - Creates the most minimal table

This demonstrates how header styling can replace traditional table borders for a modern look.
