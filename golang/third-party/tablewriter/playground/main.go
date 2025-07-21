package main

import (
	"log"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/renderer"
	"github.com/olekukonko/tablewriter/tw"
)

func main() {
	quickstart()
	table()
	streaming()
	colorizedTable()
}

func quickstart() {
	// ┌─────────────┬─────────┬────────┐
	// │   PACKAGE   │ VERSION │ STATUS │
	// ├─────────────┼─────────┼────────┤
	// │ tablewriter │ v0.0.5  │ legacy │
	// │ tablewriter │ v1.0.8  │ latest │
	// └─────────────┴─────────┴────────┘
	data := [][]string{
		{"Package", "Version", "Status"},
		{"tablewriter", "v0.0.5", "legacy"},
		{"tablewriter", "v1.0.8", "latest"},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.Header(data[0])
	_ = table.Bulk(data[1:])
	_ = table.Render()
}

func table() {
	t := tablewriter.NewTable(os.Stdout)
	t.Header([]string{"Name", "Age", "Nationality"})
	_ = t.Append([]string{"John", "25", "America"})
	_ = t.Append([]string{"Jane", "30", "England"})
	t.Footer([]string{"", "Total", "2"})
	_ = t.Render()

}

func streaming() {
	table := tablewriter.NewTable(os.Stdout, tablewriter.WithStreaming(tw.StreamConfig{Enable: true}))

	// Start streaming
	if err := table.Start(); err != nil {
		log.Fatalf("Start failed: %v", err)
	}

	defer func() { _ = table.Close() }()

	// Stream header
	table.Header([]string{"ID", "Description", "Status"})

	// Stream rows with simulated delay
	data := [][]string{
		{"1", "This description is too long", "OK"},
		{"2", "Short desc", "DONE"},
		{"3", "Another long description here", "ERROR"},
	}
	for _, row := range data {
		_ = table.Append(row)
		time.Sleep(500 * time.Millisecond) // Simulate real-time data feed
	}

	// Stream footer
	table.Footer([]string{"", "Total", "3"})
}

func colorizedTable() {
	data := [][]string{
		{"1", "This is a very long description that needs wrapping for readability", "OK"},
		{"2", "Short description", "DONE"},
		{"3", "Another lengthy description requiring truncation or wrapping", "ERROR"},
	}

	// Configure colors: green headers, cyan/magenta rows, yellow footer
	colorCfg := renderer.ColorizedConfig{
		Header: renderer.Tint{
			FG: renderer.Colors{color.FgGreen, color.Bold}, // Green bold headers
			BG: renderer.Colors{color.BgHiWhite},
		},
		Column: renderer.Tint{
			FG: renderer.Colors{color.FgCyan}, // Default cyan for rows
			Columns: []renderer.Tint{
				{FG: renderer.Colors{color.FgMagenta}}, // Magenta for column 0
				{},                                     // Inherit default (cyan)
				{FG: renderer.Colors{color.FgHiRed}},   // High-intensity red for column 2
			},
		},
		Footer: renderer.Tint{
			FG: renderer.Colors{color.FgYellow, color.Bold}, // Yellow bold footer
			Columns: []renderer.Tint{
				{},                                      // Inherit default
				{FG: renderer.Colors{color.FgHiYellow}}, // High-intensity yellow for column 1
				{},                                      // Inherit default
			},
		},
		Border:    renderer.Tint{FG: renderer.Colors{color.FgWhite}}, // White borders
		Separator: renderer.Tint{FG: renderer.Colors{color.FgWhite}}, // White separators
	}

	table := tablewriter.NewTable(os.Stdout,
		tablewriter.WithRenderer(renderer.NewColorized(colorCfg)),
		tablewriter.WithConfig(tablewriter.Config{
			Row: tw.CellConfig{
				Formatting:   tw.CellFormatting{AutoWrap: tw.WrapNormal}, // Wrap long content
				Alignment:    tw.CellAlignment{Global: tw.AlignLeft},     // Left-align rows
				ColMaxWidths: tw.CellWidth{Global: 100},
			},
			Footer: tw.CellConfig{
				Alignment: tw.CellAlignment{Global: tw.AlignRight},
			},
		}),
	)

	table.Header([]string{"ID", "Description", "Status"})
	table.Bulk(data)
	table.Footer([]string{"", "Total", "3"})
	table.Render()
}
