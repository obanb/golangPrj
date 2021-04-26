package domain

import (
	"awesomeProject/dto"
	errs "awesomeProject/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Issue struct {
	IssueId     string  `db:"issue_id" json:"issueId"`
	Name        string  `db:"name" json:"name"`
	Description string  `db:"description" json:"description"`
	CreatedAt   string  `db:"createdAt" json:"createdAt"`
	Status      float64 `db:"status" json:"status"`
	AccountId   string  `db:"account_id" json:"account_id"`
}

type IssueMongo struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string  `bson:"name,omitempty"`
	Description string  `bson:"description,omitempty"`
	CreatedAt   string 	`bson:"createdAt,omitempty"`
	Status      float64 `bson:"status,omitempty"`
	AccountId   string  `bson:"accountId,omitempty"`
}

type IssueRepository interface {
	Save(Issue) (*Issue, *errs.AppError)
	SaveMongo(Issue) (*IssueMongo, *errs.AppError)
	SaveMany([]Issue) (*[]Issue, *errs.AppError)
	FindAll() (*[]Issue, *errs.AppError)
}

func (i Issue) ToCreateIssueResponseDto() dto.CreateIssueResponse {
	return dto.CreateIssueResponse{IssueId: i.IssueId}
}

func (i Issue) ToCreateIssueResponse2Dto() *dto.CreateIssueResponse {
	return &dto.CreateIssueResponse{IssueId: i.IssueId}
}
