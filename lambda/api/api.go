package api

import (
	"lambda-func/database"
)

type ApiHandler struct {
	dbStore database.DynamoDBClient
}

func NewApiHandler(dbstore database.DynamoDBClient) ApiHandler {
	return ApiHandler{
		dbStore: dbstore,
	}
}
