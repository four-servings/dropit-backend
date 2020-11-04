package stuff

import (
	"github/four-servings/dropit-backend/setup"
	"github/four-servings/dropit-backend/stuff/api"
	"github/four-servings/dropit-backend/stuff/app/command"
	"github/four-servings/dropit-backend/stuff/infra"
	"net/http"
)

func init() {
	dbConnection := setup.GetDatabaseConnection()
	repository := infra.NewRepository(dbConnection)
	s3Adaptor := infra.NewS3Adaptor()
	commandBus := command.NewBus(repository, s3Adaptor)
	controller := api.NewController(commandBus)

	http.HandleFunc("/stuffs", controller.Handle)
}
