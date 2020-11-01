package api

import (
	"encoding/json"
	"github/four-servings/dropit-backend/stuff/app/command"
	"net/http"
)

// NewController create controller instance
func NewController(commandBus command.Bus) Controller {
	return &controllerImplement{commandBus}
}

type (
	// Controller http controller
	Controller interface {
		Handle(w http.ResponseWriter, r *http.Request)
	}

	controllerImplement struct {
		commandBus command.Bus
	}
)

// Handle handle http request
func (c *controllerImplement) Handle(w http.ResponseWriter, r *http.Request) {
	c.branchByMethod(w, r)
}

func (c *controllerImplement) branchByMethod(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		c.handlePOST(w, r)
		return
	default:
		c.handleNotAllowedMethod(w, r)
		return
	}
}

type createStuffBody struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Folder   string `json:"folder"`
}

func (c *controllerImplement) handlePOST(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	body := createStuffBody{}
	if decoder.Decode(&body) != nil {
		http.Error(w, "can not parse body", http.StatusBadRequest)
		return
	}

	name := body.Name
	category := body.Category
	folder := body.Folder

	command := &command.CreateStuff{name, category, folder}
	c.commandBus.Handle(command)
	w.WriteHeader(http.StatusCreated)
}

func (c *controllerImplement) handleNotAllowedMethod(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
}
