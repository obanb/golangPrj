package service

import (
	"awesomeProject/domain"
	errs "awesomeProject/errors"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

type XlsxService interface {
	Generate() (*[]domain.Issue, *errs.AppError)
}

func Generate() *excelize.File{
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet2")
	// Set value of a cell.
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
	return f
}