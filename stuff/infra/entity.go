package infra

import "time"

// Entity stuff entity
type Entity struct {
	ID        string    `gorm:"primary_key"`
	Name      string    `gorm:"not null"`
	Category  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}
