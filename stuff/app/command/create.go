package command

import (
	"github/four-servings/dropit-backend/stuff/domain"
	"github/four-servings/dropit-backend/stuff/infra"
)

type (
	// CreateStuff create stuff command
	CreateStuff struct {
		Name     string
		Category string
	}

	createStuffHandler interface {
		handle(command *CreateStuff)
	}

	createStuffHandlerImplement struct {
		repository infra.StuffRepository
	}
)

func newCreateStuffHandler(repository infra.StuffRepository) createStuffHandler {
	return &createStuffHandlerImplement{repository}
}

func (h *createStuffHandlerImplement) handle(command *CreateStuff) {
	id := h.repository.FindNewID()
	stuff := domain.NewStuff(id, command.Name, command.Category)
	h.repository.Save(stuff)
}
