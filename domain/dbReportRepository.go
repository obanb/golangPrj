package domain

import (
	"awesomeProject/dto"
	errs "awesomeProject/errors"
	"awesomeProject/logger"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type DbReportRepositoryCrossDb struct {
	clientMongo *mongo.Client
}

func (d DbReportRepositoryCrossDb) Save(dbr *DbReport) (*DbReport, *errs.AppError) {
	collection := d.clientMongo.Database("localhost").Collection("dbReports")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	dbReport := DbReport{
		ID:           primitive.ObjectID{},
		Name:         dbr.Name,
		Description:  dbr.Description,
		CreatedAt:    dbr.CreatedAt,
		Status:       0,
		AccountId:    dbr.AccountId,
		ReportQuery:  dbr.ReportQuery,
		ReportSource: dbr.ReportSource,
		ResultData: dbr.ResultData,
	}

	inserted, err := collection.InsertOne(ctx, dbReport)

	if err != nil {
		logger.Error("Error while creating dbReport: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	// type assertion to ObjectId primitive
	if oid, ok := inserted.InsertedID.(primitive.ObjectID); ok {
		dbReport.ID = oid
	} else {
		return nil, errs.NewUnexpectedError("Error while converting InsertedId")
	}

	if err != nil {
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	return &dbReport, nil
}

func (d DbReportRepositoryCrossDb) ExecMongoQuery(query *dto.CreateDbReportRequest) (*[]map[string]interface{}, *string, *errs.AppError) {
	collection := d.clientMongo.Database("localhost").Collection(query.Source)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var resultDataRaw []map[string]interface{}
	var tempBsonQuery bson.D

	var marshalError error

	err := bson.UnmarshalExtJSON([]byte(query.Query), true, &tempBsonQuery)
	if err != nil {
		return nil, nil, errs.NewUnexpectedError("ExecMongoQuery execution error: " + marshalError.Error())
	}

	cursor, findError := collection.Find(ctx, tempBsonQuery)
	fmt.Print("query ", tempBsonQuery)
	if findError != nil {
		return nil, nil, errs.NewUnexpectedError("ExecMongoQuery execution error: " + findError.Error())
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		t := map[string]interface{}{}
		err := cursor.Decode(&t)
		if err != nil {
			return nil, nil, errs.NewUnexpectedError("Unexpected database error")
		}
		resultDataRaw = append(resultDataRaw, t)
	}

	stringResult, stringifyError := stringifiesRawData(resultDataRaw)

	if stringifyError != nil {
		return nil, nil, errs.NewUnexpectedError("stringify error: " + stringifyError.Message)
	}

	return &resultDataRaw, stringResult, nil
}

func stringifiesRawData(rawData []map[string]interface{}) (*string, *errs.AppError) {

	b, err := json.Marshal(rawData)

	if err != nil {
		return nil, errs.NewUnexpectedError("stringifiesRawData error")
	}

	stringifies := string(b)

	return &stringifies, nil
}


func NewDbReportRepository(clientMongo *mongo.Client) DbReportRepository {
	return DbReportRepositoryCrossDb{clientMongo}
}

