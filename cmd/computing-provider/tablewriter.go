package main

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

type VisualTable struct {
	Header   []string
	Data     [][]string
	RowColor []RowColor
	WrapText bool
}

type RowColor struct {
	row    int
	column []int
	color  []tablewriter.Colors
}

func NewVisualTable(header []string, data [][]string, rowColor []RowColor) *VisualTable {

	return &VisualTable{
		Header:   header,
		Data:     data,
		RowColor: rowColor,
		WrapText: true,
	}
}

func (v *VisualTable) SetAutoWrapText(wrapText bool) *VisualTable {
	v.WrapText = wrapText
	return v
}

func (v *VisualTable) Generate(formatHeaders bool) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(v.WrapText)

	for index, datum := range v.Data {
		var rowColors []tablewriter.Colors
		for _, rowColor := range v.RowColor {
			if index == rowColor.row {
				for dIndex := range datum {
					var defaultFlag = true
					for n, colIndex := range rowColor.column {
						if dIndex == colIndex {
							rowColors = append(rowColors, rowColor.color[n])
							defaultFlag = false
						}
					}
					if defaultFlag {
						rowColors = append(rowColors, tablewriter.Colors{})
					}
				}
			}
		}
		table.Rich(v.Data[index], rowColors)
	}

	table.SetHeader(v.Header)
	table.SetAutoFormatHeaders(formatHeaders)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderLine(false)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetBorder(false)
	table.SetTablePadding("\t")
	table.SetNoWhiteSpace(true)
	table.Render()
}
