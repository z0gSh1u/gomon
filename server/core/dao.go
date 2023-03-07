package core

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	mongoOptions "go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoClient   *mongo.Client
	gomonDatabase *mongo.Database
)

const (
	DAOSuccess = 0
	DAOFailed  = 1
)

type DAOResponse struct {
	Status int
	Msg    string
}

func init() {
	var err error
	options := mongoOptions.Client().ApplyURI(MongoDBURI)

	if mongoClient, err = mongo.Connect(context.TODO(), options); err != nil {
		panic("err connect mongodb")
	}

	if err = mongoClient.Ping(context.TODO(), nil); err != nil {
		panic("err connect mongodb")
	}

	gomonDatabase = mongoClient.Database("gomon")
}

func InsertComment(comment *Comment) (r *DAOResponse) {
	if _, err := gomonDatabase.Collection("comments").InsertOne(context.TODO(), comment); err != nil {
		return &DAOResponse{Status: DAOFailed, Msg: fmt.Sprint(err)}
	}
	return &DAOResponse{Status: DAOSuccess}
}
