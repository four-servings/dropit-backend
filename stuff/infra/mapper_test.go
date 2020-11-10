package infra

import (
	"github/four-servings/dropit-backend/stuff/domain"
	"testing"
)

func TestModelToEntity(t *testing.T) {
	id := "id"
	name := "name"
	category := "category"
	folder := "folder"
	userID := "userID"
	fileKey := "image file Key"
	model := domain.NewStuff(id, name, category, folder, userID, fileKey)

	createdAt := model.CreatedAt()
	updatedAt := model.UpdatedAt()

	entity := Stuff{id, name, category, folder, userID, fileKey, createdAt, updatedAt}

	result := ModelToEntity(model)

	if result != entity {
		t.Fail()
	}
}

func TestEntityToModel(t *testing.T) {
	id := "id"
	name := "name"
	category := "category"
	folder := "folder"
	userID := "userID"
	fileKey := "image file Key"
	model := domain.NewStuff(id, name, category, folder, userID, fileKey)

	createdAt := model.CreatedAt()
	updatedAt := model.UpdatedAt()

	entity := Stuff{id, name, category, folder, userID, fileKey, createdAt, updatedAt}

	result := EntityToModel(entity)

	if result != model {
		t.Fail()
	}
}
