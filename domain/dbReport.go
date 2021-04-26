package domain

import (
	"awesomeProject/dto"
	errs "awesomeProject/errors"
)

type DbReport struct {
	ReportId     string  `json:"reportId" db:"report_id"`
	Name         string  `json:"name" db:"name"`
	Description  string  `json:"description" db:"description"`
	CreatedAt    string  `json:"createdAt" db:"created_at"`
	Status       float64 `json:"status" db:"status"`
	AccountId    string  `json:"account_id" db:"account_id"`
	IssueId      string  `json:"issueId" db:"issue_id`
	ReportQuery  string  `json:"query" db:"query"`
	ReportSource string  `json:"source" db:"report_source"`
	ResultData   *string `json:"result" db:"result_data"`
}

func (r DbReport) ToCreateDbReportResponseDto() dto.CreateDbReportResponse {
	return dto.CreateDbReportResponse{ReportId: r.ReportId}
}

type DbReportRepository interface {
	Save(DbReport) (*DbReport, *errs.AppError)
	ExecMongoQuery(query *dto.CreateDbReportRequest) (*[]map[string]interface{}, *string, *errs.AppError)
}
