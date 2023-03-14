package core

// This is the dataset access object (DAO) for MongoDB.

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
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
	Status  int
	Payload interface{}
	Msg     string
}

func init() {
	var err error

	credential := mongoOptions.Credential{
		Username: MongoAuthUsername,
		Password: MongoAuthPassword,
	}
	options := mongoOptions.Client().ApplyURI(MongoDBURI).SetAuth(credential)

	if mongoClient, err = mongo.Connect(context.TODO(), options); err != nil {
		panic("err connect mongodb " + err.Error())
	}

	if err = mongoClient.Ping(context.TODO(), nil); err != nil {
		panic("err connect mongodb " + err.Error())
	}

	gomonDatabase = mongoClient.Database("gomon")

	// defer func() {
	// 	if err := mongoClient.Disconnect(context.TODO()); err != nil {
	// 		panic("mongodb disconnect panic")
	// 	}
	// }()
}

func DAOInsertComment(comment *Comment) (r *DAOResponse) {
	if _, err := gomonDatabase.Collection("comments").InsertOne(context.TODO(), comment); err != nil {
		return &DAOResponse{Status: DAOFailed, Msg: fmt.Sprint(err)}
	}
	return &DAOResponse{Status: DAOSuccess}
}

func DAORemoveComment(comment *Comment) (r *DAOResponse) {
	if _, err := gomonDatabase.Collection("comments").DeleteOne(context.TODO(), comment); err != nil {
		return &DAOResponse{Status: DAOFailed, Msg: fmt.Sprint(err)}
	}
	return &DAOResponse{Status: DAOSuccess}
}

func DAOUpdateComment(comment *Comment) (r *DAOResponse) {
	if _, err := gomonDatabase.Collection("comments").UpdateByID(context.TODO(), bson.D{{Key: "_id", Value: comment.Id}}, comment); err != nil {
		return &DAOResponse{Status: DAOFailed, Msg: fmt.Sprint(err)}
	}
	return &DAOResponse{Status: DAOSuccess}
}

func DAOSelectCommentsByTopicId(topicId string) (r *DAOResponse) {
	cursor, err := gomonDatabase.Collection("comments").Find(context.TODO(), bson.D{{Key: "topicid", Value: topicId}})
	if err != nil {
		return &DAOResponse{Status: DAOFailed, Msg: fmt.Sprint(err)}
	}

	var results []Comment
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		return &DAOResponse{Status: DAOFailed, Msg: fmt.Sprint(err)}
	}

	return &DAOResponse{Status: DAOSuccess, Payload: results}
}
