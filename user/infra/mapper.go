package infra

import "github/four-servings/dropit-backend/domain"

// ModelToEntity convert user model to entity
func ModelToEntity(model domain.User) UserEntity {
	id := model.ID()
	deviceID := model.DeviceID()
	createdAt := model.CreatedAt()
	updatedAt := model.UpdatedAt()
	deletedAt := model.DeletedAt()
	return UserEntity{id, deviceID, createdAt, updatedAt, deletedAt}
}

// EntityToModel convert user entity to model
func EntityToModel(entity UserEntity) domain.User {
	id := entity.ID
	deviceID := entity.DeviceID
	createdAt := entity.CreatedAt
	updatedAt := entity.UpdatedAt
	deletedAt := entity.DeletedAt
	anemic := domain.AnemicUser{id, deviceID, createdAt, updatedAt, deletedAt}
	return anemic.ToRichModel()
}
