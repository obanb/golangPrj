package domain

import (
	errs "awesomeProject/errors"
	"awesomeProject/logger"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
	"time"
)

type IssueRepositorySql struct {
	client *sqlx.DB
	clientMongo *mongo.Client
}

func (d IssueRepositorySql) FindAll() (*[]Issue, *errs.AppError) {
	//var rows *sql.Rows
	var err error
	issues := make([]Issue, 0)

	findAllSql := "select * from issues"
	err = d.client.Select(&issues, findAllSql)

	if err != nil {
		logger.Error("Error while quering issues from database " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	//err = sqlx.StructScan(rows, &customers)

	return &issues, nil
}

func (d IssueRepositorySql) Save(i Issue) (*Issue, *errs.AppError) {
	sqlInsert := "INSERT INTO issues (name, description, createdAt, status, account_id) values (?,?,?,?,?)"

	result, err := d.client.Exec(sqlInsert, i.Name, i.Description, i.CreatedAt, i.Status, i.AccountId)

	if err != nil {
		logger.Error("Error while creating issue: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()

	i.IssueId = strconv.FormatInt(id, 10)
	return &i, nil
}

func (d IssueRepositorySql) SaveMongo(i Issue) (*IssueMongo, *errs.AppError) {
	collection := d.clientMongo.Database("localhost").Collection("issues")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	issue := IssueMongo{
		Name:        i.Name,
		Description: i.Description,
		CreatedAt:   i.CreatedAt,
		Status:      0,
		AccountId:   i.AccountId,
	}

	_, err := collection.InsertOne(ctx, issue)

	if err != nil {
		logger.Error("Error while creating issue: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
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

func NewIssueRepositorySql(dbClient *sqlx.DB, clientMongo *mongo.Client) IssueRepositorySql {
	return IssueRepositorySql{dbClient,clientMongo }
}
