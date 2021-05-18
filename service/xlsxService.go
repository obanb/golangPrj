package service

import (
	errs "awesomeProject/errors"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"time"
)

type XlsxService interface {
	Generate() (*excelize.File, *string, *errs.AppError)
}

type DefaultXlsxService struct {
}

func (s DefaultXlsxService) Generate() (*excelize.File, *string, *errs.AppError) {
	f := excelize.NewFile()
	indexName := "report"

	index := f.NewSheet(indexName)

	x := [5]string{"pes","kocka","vlocka","prase","zase"}

	for  i,value :=  range x {
		cell, err := excelize.CoordinatesToCellName(i + 1,1)
		if err != nil{
			appError := errs.NewUnexpectedError(err.Error())
			return nil, nil, appError
		}
		f.SetCellValue(indexName, cell, value)
	}

	f.SetActiveSheet(index)

	filename := time.Now().UTC().Format("report-20060102150405.xlsx")

	return f, &filename, nil
}

func NewDbXlsxService() DefaultXlsxService {
	return DefaultXlsxService{}
}
