package infra

import (
	"errors"
	"github/four-servings/dropit-backend/user/config"
	"github/four-servings/dropit-backend/user/domain"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type (
	// UserRepository user repository interface
	UserRepository interface {
		FindNewID() string
		Save(user domain.User)
	}

	// userRepositoryImplements user repository implements
	userRepositoryImplements struct {
		db *gorm.DB
	}
)

// NewRepository create repository instance
func NewRepository() UserRepository {
	db := getDatabaseConnection()
	db.AutoMigrate(&Entity{})
	return &userRepositoryImplements{db: db}
}

// FindNewID find new entity id
func (r *userRepositoryImplements) FindNewID() string {
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
func (r *userRepositoryImplements) Save(user domain.User) {
	entity := ModelToEntity(user)
	if err := r.db.Save(entity).Error; err != nil {
		panic(nil)
	}
}

func getDatabaseConnection() *gorm.DB {
	dbConfig := config.Database{}

	user := dbConfig.User()
	password := dbConfig.Password()
	host := dbConfig.Host()
	port := dbConfig.Port()
	name := dbConfig.Name()

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + name + "?parseTime=true"

	connection, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{
			Logger: logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags),
				logger.Config{
					SlowThreshold: time.Second,
					LogLevel:      logger.Silent,
				},
			),
		},
	)
	if err != nil {
		panic(err)
	}

	return connection
}
