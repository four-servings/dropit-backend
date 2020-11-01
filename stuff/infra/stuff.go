package infra

import "time"

// Stuff stuff entity
type Stuff struct {
	ID        string    `gorm:"primary_key"`
	Name      string    `gorm:"not null"`
	Category  string    `gorm:"not null"`
	Folder    string    `gorm:"not null"`
	UserID    string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}
