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
		Folder   string
		UserID   string
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
	name := command.Name
	category := command.Category
	folder := command.Folder
	userID := command.UserID
	stuff := domain.NewStuff(id, name, category, folder, userID)
	h.repository.Save(stuff)
}
