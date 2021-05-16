package domain

import (
	errs "awesomeProject/errors"
	"awesomeProject/logger"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type IssueRepositorySql struct {
	clientMongo *mongo.Client
}

func (d IssueRepositorySql) FindAll() (*[]Issue, *errs.AppError) {
	issues := make([]Issue, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	collection := d.clientMongo.Database("localhost").Collection("issues")

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"_id", -1}})

	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)

	defer cur.Close(ctx)

	if err != nil {
		logger.Error("Error while quering issues from database " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	for cur.Next(context.TODO()) {
		t := Issue{}
		err := cur.Decode(&t)
		if err != nil {
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
		issues = append(issues, t)
	}

	return &issues, nil
}

func (d IssueRepositorySql) Save(i Issue) (*Issue, *errs.AppError) {
	collection := d.clientMongo.Database("localhost").Collection("issues")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	issue := Issue{
		Name:        i.Name,
		Description: i.Description,
		CreatedAt:   i.CreatedAt,
		Status:      0,
		AccountId:   i.AccountId,
	}

	inserted, err := collection.InsertOne(ctx, issue)

	if err != nil {
		logger.Error("Error while creating issue: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	// type assertion to ObjectId primitive
	if oid, ok := inserted.InsertedID.(primitive.ObjectID); ok {
		issue.ID = oid
	} else {
		return nil, errs.NewUnexpectedError("Error while converting InsertedId")
	}

	if err != nil {
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	return &issue, nil
}

func (d IssueRepositorySql) SaveMany(i []Issue) (*[]Issue, *errs.AppError) {
	results := make([]Issue, len(i))
	c := make(chan Issue, len(i))

	for _, issue := range i {
		go func(issue Issue) {
			i, err := d.Save(issue)
			if err != nil {
				fmt.Print("one issue error")
				c <- *i
			} else {
				fmt.Print("one issue error")
				c <- *i
			}
		}(issue)
	}

	for range i {
		results = append(results, <-c)
	}
	return &results, nil
}

func NewIssueRepositorySql(clientMongo *mongo.Client) IssueRepositorySql {
	return IssueRepositorySql{clientMongo}
}
