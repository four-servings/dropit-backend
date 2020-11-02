package infra

import (
	"github/four-servings/dropit-backend/user/domain"
	"testing"
)

func TestModelToEntity(t *testing.T) {
	id := "id"
	deviceID := "deviceID"
	model := domain.NewUser(id, deviceID)

	createdAt := model.CreatedAt()
	updatedAt := model.UpdatedAt()

	entity := User{id, deviceID, createdAt, updatedAt, nil}

	result := ModelToEntity(model)

	if result != entity {
		t.Fail()
	}
}

func TestEntityToModel(t *testing.T) {
	id := "id"
	deviceID := "deviceID"
	model := domain.NewUser(id, deviceID)

	createdAt := model.CreatedAt()
	updatedAt := model.UpdatedAt()

	entity := User{id, deviceID, createdAt, updatedAt, nil}

	result := EntityToModel(entity)

	if result != model {
		t.Fail()
	}
}
