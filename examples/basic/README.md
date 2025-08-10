# Basic Table Example

This example demonstrates the basic usage of termhyo for creating simple tables.

## Running the Example

```bash
go run main.go
```

Or build and run:

```bash
go build -o basic .
./basic
```

## What This Example Shows

- Creating a simple table with columns
- Adding rows to the table
- Basic table rendering
- Column alignment (left, right, center)
- Auto-width calculation

## Code Overview

The example creates a table with ID, Name, Age, and City columns and demonstrates:

1. **Column Definition**: Setting up columns with different alignments
2. **Table Creation**: Using `termhyo.NewTable()`
3. **Adding Data**: Using `table.AddRow()` to add data
4. **Rendering**: Using `table.Render()` to output the table

This is the most fundamental example and a good starting point for understanding termhyo.
