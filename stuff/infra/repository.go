package infra

import (
	"github/four-servings/dropit-backend/stuff/domain"

	"gorm.io/gorm"
)

type (
	// StuffRepository stuff repository interface
	StuffRepository interface {
		Save(stuff domain.Stuff)
	}

	stuffRepositoryImplement struct {
		db *gorm.DB
	}
)

// NewRepository create repository instance
func NewRepository(db *gorm.DB) StuffRepository {
	db.AutoMigrate(&Entity{})
	return &stuffRepositoryImplement{db}
}

// Save insert or update stuff data
func (r *stuffRepositoryImplement) Save(stuff domain.Stuff) {
	entity := ModelToEntity(stuff)
	if err := r.db.Save(entity).Error; err != nil {
		panic(err)
	}
}
