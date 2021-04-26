package service

import (
	"awesomeProject/domain"
	"awesomeProject/dto"
	errs "awesomeProject/errors"
	"time"
)

type DbReportService interface {
	CreateDbReport(request *dto.CreateDbReportRequest) (*dto.CreateDbReportResponse, *errs.AppError)
}

type DefaultDbReportService struct {
	repo domain.DbReportRepository
}

func (s DefaultDbReportService) CreateDbReport(req *dto.CreateDbReportRequest) (*dto.CreateDbReportResponse, *errs.AppError) {
	DbReport := domain.DbReport{
		ReportId:     "1",
		Name:         req.Name,
		Description:  req.Description,
		CreatedAt:    time.Now().Format("2006-01-02 15:04:05"),
		AccountId:    "95470",
		Status:       1,
		ReportQuery:  req.Query,
		ReportSource: req.Source,
		IssueId:      "2006",
	}

	_, stringResult, err := s.repo.ExecMongoQuery(req)

	if err != nil {
		return nil, errs.NewUnexpectedError("CreateDbReport ExecMongoQuery error " + err.Message)
	}

	DbReport.ResultData = stringResult

	newDbReport, err := s.repo.Save(DbReport)
	if err != nil {
		return nil, err
	}

	response := newDbReport.ToCreateDbReportResponseDto()

	return &response, nil
}

func NewDbReportService(repository domain.DbReportRepository) DefaultDbReportService {
	return DefaultDbReportService{repository}
}
