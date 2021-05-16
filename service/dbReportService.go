package service

import (
	"awesomeProject/domain"
	"awesomeProject/dto"
	errs "awesomeProject/errors"
	"fmt"
	"time"
)

type DbReportService interface {
	CreateDbReport(request *dto.CreateDbReportRequest) (*domain.DbReport, *errs.AppError)
}

type DefaultDbReportService struct {
	repo domain.DbReportRepository
}

func (s DefaultDbReportService) CreateDbReport(req *dto.CreateDbReportRequest) (*domain.DbReport, *errs.AppError) {
	DbReport := &domain.DbReport{
		Name:         req.Name,
		Description:  req.Description,
		CreatedAt:    time.Now().Format("2006-01-02 15:04:05"),
		AccountId:    "95470",
		Status:       1,
		ReportQuery:  req.Query,
		ReportSource: req.Source,
	}

	_, stringResult, err := s.repo.ExecMongoQuery(req)

	if err != nil {
		return nil, errs.NewUnexpectedError("CreateDbReport ExecMongoQuery error " + err.Message)
	}

	DbReport.ResultData = *stringResult

	newDbReport, err := s.repo.Save(DbReport)
	if err != nil {
		return nil, err
	}

	fmt.Print(newDbReport)

	return newDbReport, nil
}

func NewDbReportService(repository domain.DbReportRepository) DefaultDbReportService {
	return DefaultDbReportService{repository}
}
