package domain

import (
	"awesomeProject/dto"
	errs "awesomeProject/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DbReport struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty`
	Name         string             `json:"name" bson:"name"`
	Description  string             `json:"description" bson:"description"`
	CreatedAt    string             `json:"createdAt" bson:"created_at"`
	Status       float64            `json:"status" bson:"status"`
	AccountId    string             `json:"account_id" bson:"account_id"`
	ReportQuery  string             `json:"query" bson:"query"`
	ReportSource string             `json:"source" bson:"report_source"`
	ResultData   string             `json:"result, omitempty" bson:"result_data,omitempty"`
}

type DbReportRepository interface {
	Save(*DbReport) (*DbReport, *errs.AppError)

	ExecMongoQuery(query *dto.CreateDbReportRequest) (*[]map[string]interface{}, *string, *errs.AppError)
}
