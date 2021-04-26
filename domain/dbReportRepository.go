package domain

import (
	"awesomeProject/dto"
	errs "awesomeProject/errors"
	"awesomeProject/logger"
	"context"
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
	"time"
)

type DbReportRepositoryCrossDb struct {
	clientSql   *sqlx.DB
	clientMongo *mongo.Client
}

func (d DbReportRepositoryCrossDb) Save(dbr DbReport) (*DbReport, *errs.AppError) {
	sqlInsert := "INSERT INTO dbReports (name, description, created_At, status, account_id, issue_id, report_query, report_source, result_data) values (?,?,?,?,?,?,?,?,?)"

	result, err := d.clientSql.Exec(sqlInsert, dbr.Name, dbr.Description, dbr.CreatedAt, dbr.Status, dbr.AccountId, dbr.IssueId, dbr.ReportQuery, dbr.ReportSource, dbr.ResultData)

	if err != nil {
		logger.Error("Error while creating DbReport: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()

	dbr.ReportId = strconv.FormatInt(id, 10)
	return &dbr, nil
}

func (d DbReportRepositoryCrossDb) ExecMongoQuery(query *dto.CreateDbReportRequest) (*[]map[string]interface{}, *string, *errs.AppError) {
	collection := d.clientMongo.Database("localhost").Collection(query.Source)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resultDataRaw := make([]map[string]interface{}, 0)
	item := map[string]interface{}{}
	var tempBsonDocument bson.D
	var tempBsonQuery bson.D
	var tempBytes []byte

	var marshalError error

	marshalError = bson.UnmarshalExtJSON([]byte(query.Query), true, &tempBsonQuery)
	if marshalError != nil {
		return nil, nil, errs.NewUnexpectedError("ExecMongoQuery execution error: " + marshalError.Error())
	}

	cursor, findError := collection.Find(ctx, tempBsonQuery)
	if findError != nil {
		return nil, nil, errs.NewUnexpectedError("ExecMongoQuery execution error: " + findError.Error())
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		// set to bsonDocument temporary
		decodeError := cursor.Decode(&tempBsonDocument)
		if decodeError != nil {
			return nil, nil, errs.NewUnexpectedError("cursor decode error: " + findError.Error())
		}

		tempBytes, marshalError = bson.Marshal(tempBsonDocument)
		if marshalError != nil {
			return nil, nil, errs.NewUnexpectedError("marshal error: " + marshalError.Error())
		}

		// unmarshal temporary bytes to map
		marshalError = bson.Unmarshal(tempBytes, item)
		if marshalError != nil {
			return nil, nil, errs.NewUnexpectedError("unmarshal error: " + marshalError.Error())
		}

		// append item to unknown interface
		resultDataRaw = append(resultDataRaw, item)
	}

	stringResult, stringifyError := stringifiesRawData(resultDataRaw)

	if stringifyError != nil {
		return nil, nil, errs.NewUnexpectedError("stringify error: " + stringifyError.Message)
	}

	return &resultDataRaw, stringResult, nil
}

func NewDbReportRepository(clientSql *sqlx.DB, clientMongo *mongo.Client) DbReportRepository {
	return DbReportRepositoryCrossDb{clientSql, clientMongo}
}

func stringifiesRawData(rawData []map[string]interface{}) (*string, *errs.AppError) {

	b, err := json.Marshal(rawData)

	if err != nil {
		return nil, errs.NewUnexpectedError("stringifiesRawData error")
	}

	stringifies := string(b)

	return &stringifies, nil
}
