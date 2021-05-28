package domain

import (
	errs "awesomeProject/errors"
	"go.mongodb.org/mongo-driver/bson"
)

type FileMetadata struct {
	FsId string
	Filename string
	Date string
	Len *string
}

func (fm FileMetadata) ToBson() (*bson.D, *errs.AppError) {
	var tempBson bson.D
	data, err := bson.Marshal(fm)

	if err!= nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	err = bson.Unmarshal(data, &tempBson)

	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	return &tempBson, nil
}
