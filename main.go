package main

import (
	"github/four-servings/dropit-backend/setup"
	"github/four-servings/dropit-backend/user/api"
	"github/four-servings/dropit-backend/user/app/command"
	"github/four-servings/dropit-backend/user/infra"
	"net/http"
)

func main() {
	dbConnection := setup.GetDatabaseConnection()
	repository := infra.NewRepository(dbConnection)
	commandBus := command.NewBus(repository)
	controller := api.NewController(commandBus)

	http.HandleFunc("/users", controller.Handle)

	http.ListenAndServe(":5000", nil)
}
