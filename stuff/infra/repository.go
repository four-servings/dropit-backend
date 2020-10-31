package infra

import (
	"errors"
	"github/four-servings/dropit-backend/stuff/domain"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type (
	// StuffRepository stuff repository interface
	StuffRepository interface {
		Save(stuff domain.Stuff)
		FindNewID() string
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

// FindNewID find new entity id
func (r *stuffRepositoryImplement) FindNewID() string {
	uuid, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}
	result := r.db.Where(&Entity{ID: uuid.String()}).First(&Entity{})

	if result.Error != nil && result.Error.Error() != "record not found" {
		panic(result.Error)
	}

	if result.RowsAffected != 0 {
		panic(errors.New("found id is duplicated"))
	}

	return uuid.String()
}
