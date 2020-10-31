package infra

import "time"

// UserEntity user entity
type UserEntity struct {
	ID        string `gorm:"primary_key"`
	DeviceID  string `gorm:"unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
