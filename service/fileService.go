package service

import (
	"awesomeProject/domain"
	errs "awesomeProject/errors"
	"bytes"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"strconv"
	"time"
)

type FileService interface {
	Save(buffer []byte, metadata *domain.FileMetadata) (*string, *errs.AppError)
	DownloadById(id string) (*bytes.Buffer, *errs.AppError)
}

type DefaultFileService struct {
	clientMongo *mongo.Client
}

func (s DefaultFileService) DownloadById(id string) (*bytes.Buffer, *errs.AppError) {
	database := s.clientMongo.Database("localhost")
	fsFiles := database.Collection("fs.files")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var results bson.M

	err := fsFiles.FindOne(ctx, bson.M{"metadata.fsid": id}).Decode(&results)

	fmt.Print(results)

	if err != nil {
		appError := errs.NewUnexpectedError(err.Error())
		return nil, appError
	}

	bucket, _ := gridfs.NewBucket(
		database,
	)

	var buf bytes.Buffer

	_, streamErr := bucket.DownloadToStreamByName("egd", &buf)

	if streamErr != nil {
		appError := errs.NewUnexpectedError(err.Error())
		return nil, appError
	}

	return &buf, nil
}

func (s DefaultFileService) Save(buffer []byte, metadata *domain.FileMetadata) (*string, *errs.AppError) {
	database := s.clientMongo.Database("localhost")
	bucket, err := gridfs.NewBucket(
		database,
	)
	if err != nil {
		appError := errs.NewUnexpectedError(err.Error())
		return nil, appError
	}

	bsonMedata, appErr := metadata.ToBson()

	if appErr != nil {
		return nil, appErr
	}

	uploadStream, err := bucket.OpenUploadStream(
		metadata.Filename,
		&options.UploadOptions{
			Metadata: bsonMedata,
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
