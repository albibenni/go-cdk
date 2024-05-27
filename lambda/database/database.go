package database

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamoDBClient struct {
	databaseStore *dynamodb.DynamoDB
}

func NewDynamoDBClient() DynamoDBClient {
	dbsession := session.Must(session.NewSession())

	db := dynamodb.New(dbsession)
	return DynamoDBClient{
		databaseStore: db,
	}
}
