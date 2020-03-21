package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

type row struct {
	leftCol  string
	rightCol string
}

type table struct {
	rows []row
}

func generatePdf(tables []table) {
	pdf := gofpdf.New("L", "mm", "A4", "")

	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	for idx, tableObj := range tables {
		pos := float64((90 * idx) + 10)

		if idx != 0 {
			pos += 10
		}

		pdf.Rect(pos, 10, 90, 80, "")
		pdf.SetY(10)

		for _, row := range tableObj.rows {
			pdf.SetX(pos)
			pdf.CellFormat(45, 10, row.leftCol, "1", 0, "L", false, 0, "")
			pdf.CellFormat(45, 10, row.rightCol, "1", 0, "L", false, 0, "")
			pdf.Ln(-1)
		}
	}

	_ = pdf.OutputFileAndClose("output/hello.pdf")
}

func main() {
	boxes := []table{
		table{
			[]row{
				row{
					leftCol:  "first",
					rightCol: "second",
				},
				row{
					leftCol:  "third",
					rightCol: "forth",
				},
				row{
					leftCol: "more", rightCol: "data",
				},
			},
		},
		table{
			[]row{
				row{
					leftCol:  "x",
					rightCol: "y",
				},
				row{
					leftCol:  "z",
					rightCol: "q",
				},
			},
		},
	}

	generatePdf(boxes)
	fmt.Println("success")
}
