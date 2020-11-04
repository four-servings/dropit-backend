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
	UserID   string `json:"user_id"`
}

func (c *controllerImplement) handlePOST(w http.ResponseWriter, r *http.Request) {

	_, fileHeader, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "can not found image", http.StatusBadRequest)
		return
	}

	category := r.FormValue("category")
	if category == "" {
		http.Error(w, "category is empty", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "name is empty", http.StatusBadRequest)
		return
	}

	folder := r.FormValue("folder")
	if folder == "" {
		http.Error(w, "folder is empty", http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	userID := "userID"

	command := &command.CreateStuff{name, category, folder, userID, fileHeader}
	c.commandBus.Handle(command)
	w.WriteHeader(http.StatusCreated)
}

func (c *controllerImplement) handleNotAllowedMethod(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
}
