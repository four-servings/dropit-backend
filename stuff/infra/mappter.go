package infra

import "github/four-servings/dropit-backend/stuff/domain"

// ModelToEntity convert stuff model to entity
func ModelToEntity(model domain.Stuff) Entity {
	id := model.ID()
	name := model.Name()
	category := model.Category()
	createdAt := model.CreatedAt()
	updatedAt := model.UpdatedAt()
	return Entity{id, name, category, createdAt, updatedAt}
}

// EntityToModel convert stuff entity to model
func EntityToModel(entity Entity) domain.Stuff {
	id := entity.ID
	name := entity.Name
	category := entity.Category
	createdAt := entity.CreatedAt
	updatedAt := entity.UpdatedAt
	anemic := domain.AnemicStuff{id, name, category, createdAt, updatedAt}
	return anemic.ToRichModel()
}
