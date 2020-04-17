package main

import (
	"fmt"
	"os"

	"github.com/tealeg/xlsx"
)

func main() {
	xlsxFile := xlsx.NewFile()
	sheet, err := xlsxFile.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf("add sheet fail, %v", err)
		os.Exit(1)
	}

	style := xlsx.NewStyle()

	alignment := xlsx.Alignment{
		Horizontal: "left",
		Vertical:   "center",
	}
	fill := xlsx.NewFill("solid", "00FFFFFF", "00FF0000")
	font := xlsx.NewFont(12, "Arial")
	font.Bold = true
	border := xlsx.NewBorder("thin", "thin", "thin", "thin")

	style.Alignment = alignment
	style.Fill = *fill
	style.Font = *font
	style.Border = *border

	style.ApplyAlignment = true
	style.ApplyFill = true
	style.ApplyFont = true
	style.ApplyBorder = true

	row := sheet.AddRow()
	row.SetHeightCM(1)

	cell := row.AddCell()
	cell.Value = "姓名"
	cell.SetStyle(style)

	cell = row.AddCell()
	cell.Value = "年龄"
	cell.SetStyle(style)

	cell = row.AddCell()
	cell.Value = "籍贯"
	cell.SetStyle(style)

	cell = row.AddCell()
	cell.Value = "家庭住址"
	cell.SetStyle(style)

	cell = row.AddCell()
	cell.Value = "学历"
	cell.SetStyle(style)

	row1 := sheet.AddRow()
	cell = row1.AddCell()
	cell.Value = "张三"
	cell = row1.AddCell()
	cell.Value = "25"
	cell = row1.AddCell()
	cell.Value = "杭州"
	cell = row1.AddCell()
	cell.Value = "浙江省西湖区"
	cell = row1.AddCell()
	cell.Value = "本科"

	if err := xlsxFile.Save("test.xlsx"); err != nil {
		fmt.Printf("save xlsx file fail: %v", err)
		os.Exit(1)
	}
}
