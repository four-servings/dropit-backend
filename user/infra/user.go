package infra

import "time"

// User user entity
type User struct {
	ID        string     `gorm:"primary_key"`
	DeviceID  string     `gorm:"unique;not null"`
	CreatedAt time.Time  `gorm:"not null"`
	UpdatedAt time.Time  `gorm:"not null"`
	DeletedAt *time.Time `sql:"index"`
}
