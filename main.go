package main

import (
	"github/four-servings/dropit-backend/api"
	"github/four-servings/dropit-backend/app"
	"github/four-servings/dropit-backend/infra"
	"net/http"
)

func main() {
	repository := infra.NewRepository()
	commandBus := app.NewCommandBus(repository)
	controller := api.NewController(commandBus)

	http.HandleFunc("/users", controller.Handle)

	http.ListenAndServe(":5000", nil)
}
