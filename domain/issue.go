package domain

import (
	errs "awesomeProject/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Issue struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string  `bson:"name,omitempty"`
	Description string  `bson:"description,omitempty"`
	CreatedAt   string 	`bson:"createdAt,omitempty"`
	Status      float64 `bson:"status,omitempty"`
	AccountId   string  `bson:"accountId,omitempty"`
}

type IssueRepository interface {
	Save(Issue) (*Issue, *errs.AppError)
	SaveMany([]Issue) (*[]Issue, *errs.AppError)
	FindAll() (*[]Issue, *errs.AppError)
}
