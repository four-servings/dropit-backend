package command

import (
	"github/four-servings/dropit-backend/user/domain"
	"github/four-servings/dropit-backend/user/infra"
)

type (
	// CreateUser create user command
	CreateUser struct {
		DeviceID string
	}

	// createUserHandler handler interface for CreateUser command
	createUserHandler interface {
		handle(command *CreateUser)
	}

	// createUserHandlerImplement handler implement for CreateUser command
	createUserHandlerImplement struct {
		repository infra.UserRepository
	}
)

func newCreateUserHandler(repository infra.UserRepository) createUserHandler {
	return &createUserHandlerImplement{repository}
}

func (h *createUserHandlerImplement) handle(command *CreateUser) {
	id := h.repository.FindNewID()
	user := domain.NewUser(id, command.DeviceID)
	h.repository.Save(user)
}
