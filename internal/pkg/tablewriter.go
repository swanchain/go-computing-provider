package pkg

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

type VisualTable struct {
	Header   []string
	Data     [][]string
	RowColor []RowColor
}

type RowColor struct {
	Row    int
	Column []int
	Color  []tablewriter.Colors
}

func NewVisualTable(header []string, data [][]string, rowColor []RowColor) *VisualTable {

	return &VisualTable{
		Header:   header,
		Data:     data,
		RowColor: rowColor,
	}
}

func (v *VisualTable) Generate(formatHeaders bool) {
	table := tablewriter.NewWriter(os.Stdout)

	for index, datum := range v.Data {
		var rowColors []tablewriter.Colors
		for _, rowColor := range v.RowColor {
			if index == rowColor.Row {
				for dIndex := range datum {
					var defaultFlag = true
					for n, colIndex := range rowColor.Column {
						if dIndex == colIndex {
							rowColors = append(rowColors, rowColor.Color[n])
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
	table.SetAutoWrapText(false)
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
