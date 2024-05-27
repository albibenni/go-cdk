package app

import (
	"lambda-func/api"
	"lambda-func/database"
)

type App struct {
	ApiHandler api.ApiHandler
}

func NewApp() App {

	// initialize our db and pass down into ApiHandler
	db := database.NewDynamoDBClient()
	apiHandler := api.NewApiHandler(db)
	return App{
        ApiHandler: apiHandler,
    }
}
