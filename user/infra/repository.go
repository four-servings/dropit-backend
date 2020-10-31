package infra

import (
	"errors"
	"github/four-servings/dropit-backend/user/domain"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type (
	// UserRepository user repository interface
	UserRepository interface {
		FindNewID() string
		Save(user domain.User)
	}

	userRepositoryImplement struct {
		db *gorm.DB
	}
)

// NewRepository create repository instance
func NewRepository(db *gorm.DB) UserRepository {
	db.AutoMigrate(&Entity{})
	return &userRepositoryImplement{db}
}

// FindNewID find new entity id
func (r *userRepositoryImplement) FindNewID() string {
	id, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}

	result := r.db.Where(&Entity{ID: id.String()}).First(&Entity{})

	if result.Error != nil && result.Error.Error() != "record not found" {
		panic(result.Error)
	}

	if result.RowsAffected != 0 {
		panic(errors.New("found id is duplicated"))
	}

	return id.String()
}

// Save insert or update user data
func (r *userRepositoryImplement) Save(user domain.User) {
	entity := ModelToEntity(user)
	if err := r.db.Save(entity).Error; err != nil {
		panic(nil)
	}
}
