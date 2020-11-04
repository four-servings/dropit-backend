package command

import (
	"github/four-servings/dropit-backend/stuff/domain"
	"github/four-servings/dropit-backend/stuff/infra"
	"mime/multipart"
)

type (
	// CreateStuff create stuff command
	CreateStuff struct {
		Name     string
		Category string
		Folder   string
		UserID   string
		Image    *multipart.FileHeader
	}

	createStuffHandler interface {
		handle(command *CreateStuff)
	}

	createStuffHandlerImplement struct {
		repository infra.StuffRepository
		s3Adaptor  infra.S3Adaptor
	}
)

func newCreateStuffHandler(repository infra.StuffRepository, s3Adaptor infra.S3Adaptor) createStuffHandler {
	return &createStuffHandlerImplement{repository, s3Adaptor}
}

func (h *createStuffHandlerImplement) handle(command *CreateStuff) {
	id := h.repository.FindNewID()
	name := command.Name
	category := command.Category
	folder := command.Folder
	userID := command.UserID
	fileKey := h.s3Adaptor.Upload(command.Image)
	stuff := domain.NewStuff(id, name, category, folder, userID, fileKey)
	h.repository.Save(stuff)
}
