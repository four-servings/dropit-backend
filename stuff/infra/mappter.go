package infra

import "github/four-servings/dropit-backend/stuff/domain"

// ModelToEntity convert stuff model to entity
func ModelToEntity(model domain.Stuff) Stuff {
	id := model.ID()
	name := model.Name()
	category := model.Category()
	createdAt := model.CreatedAt()
	updatedAt := model.UpdatedAt()
	return Stuff{id, name, category, createdAt, updatedAt}
}

// EntityToModel convert stuff entity to model
func EntityToModel(entity Stuff) domain.Stuff {
	id := entity.ID
	name := entity.Name
	category := entity.Category
	createdAt := entity.CreatedAt
	updatedAt := entity.UpdatedAt
	anemic := domain.AnemicStuff{id, name, category, createdAt, updatedAt}
	return anemic.ToRichModel()
}
