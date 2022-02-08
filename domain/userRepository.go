package domain

import (
	errs "awesomeProject/errors"
	"awesomeProject/logger"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type UserRepositoryMongo struct {
	clientMongo *mongo.Client
}

func (d UserRepositoryMongo) Save(u User) (*User, *errs.AppError) {
	collection := d.clientMongo.Database("localhost").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	inserted, err := collection.InsertOne(ctx, u)

	if err != nil {
		logger.Error("Error while creating issue: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	// type assertion to ObjectId primitive
	if oid, ok := inserted.InsertedID.(primitive.ObjectID); ok {
		u.ID = oid
	} else {
		return nil, errs.NewUnexpectedError("Error while converting InsertedId")
	}

	if err != nil {
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	return &u, nil
}

func NewUserRepository(clientMongo *mongo.Client) UserRepositoryMongo {
	return UserRepositoryMongo{clientMongo}
}
