package main

import (
	"github/four-servings/dropit-backend/user/api"
	"github/four-servings/dropit-backend/user/app"
	"github/four-servings/dropit-backend/user/infra"
	"net/http"
)

func main() {
	repository := infra.NewRepository()
	commandBus := app.NewCommandBus(repository)
	controller := api.NewController(commandBus)

	http.HandleFunc("/users", controller.Handle)

	http.ListenAndServe(":5000", nil)
}
