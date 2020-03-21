package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

func generatePdf(boxes [][][2]string) {
	pdf := gofpdf.New("L", "mm", "A4", "")

	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	for idx, box := range boxes {
		pos := float64((90 * idx) + 10)

		if idx != 0 {
			pos += 10
		}

		pdf.Rect(pos, 10, 90, 80, "")
		pdf.SetY(10)

		for _, val := range box {
			pdf.SetX(pos)
			pdf.CellFormat(45, 10, val[0], "1", 0, "L", false, 0, "")
			pdf.CellFormat(45, 10, val[1], "1", 0, "L", false, 0, "")
			pdf.Ln(-1)
		}
	}

	_ = pdf.OutputFileAndClose("output/hello.pdf")
}

func main() {
	values := [][][2]string{
		[][2]string{
			[2]string{"first", "second"},
			[2]string{"last", "whatevs"},
		},
		[][2]string{
			[2]string{"third", "fourth"},
			[2]string{"MOAR", "valuez"},
		},
	}

	generatePdf(values)
	fmt.Println("success")
}
