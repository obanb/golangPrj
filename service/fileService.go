package service

import (
	errs "awesomeProject/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
)

type FileService interface {
	Save(buffer []byte) (*string, *errs.AppError)
}

type DefaultFileService struct {
	clientMongo *mongo.Client
}


type GridFsMetadata struct {
	fsId string
}

func (s DefaultFileService) Save(buffer []byte) (*string, *errs.AppError)  {
	database := s.clientMongo.Database("localhost")
	bucket, err := gridfs.NewBucket(
		database,
	)
	if err!= nil {
		appError := errs.NewUnexpectedError(err.Error())
		return nil, appError
	}

	uploadStream, err := bucket.OpenUploadStream(
		"test",
		&options.UploadOptions{
			Metadata: bson.D{{"fileSystemId", "egeggegge"}},
		},
	)

	defer uploadStream.Close()

	if err != nil {
		appError := errs.NewUnexpectedError(err.Error())
		return nil, appError
	}

	fileSize, err := uploadStream.Write(buffer)

	if err != nil {
		appError := errs.NewUnexpectedError(err.Error())
		return nil, appError
	}

	res := strconv.Itoa(fileSize)
	return &res, nil
}

func NewFileService(clientMongo *mongo.Client) DefaultFileService {
	return DefaultFileService{clientMongo}
}
