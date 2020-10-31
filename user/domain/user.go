package domain

import (
	"time"
)

type (
	// User user model
	User struct {
		id        string
		deviceID  string
		createdAt time.Time
		updatedAt time.Time
		deletedAt *time.Time
	}

	// AnemicUser anemic user model
	AnemicUser struct {
		ID        string
		DeviceID  string
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt *time.Time
	}
)

// ID get user id
func (user *User) ID() string {
	return user.id
}

// DeviceID get user deviceID
func (user *User) DeviceID() string {
	return user.deviceID
}

// CreatedAt get user createdAt
func (user *User) CreatedAt() time.Time {
	return user.createdAt
}

// UpdatedAt get user updatedAt
func (user *User) UpdatedAt() time.Time {
	return user.updatedAt
}

// DeletedAt get user deletedAt
func (user *User) DeletedAt() *time.Time {
	return user.deletedAt
}

// NewUser create new user
func NewUser(id, deviceID string) User {
	return User{id, deviceID, time.Now(), time.Now(), nil}
}

// ToRichModel create rich user model from anemic model
func (a *AnemicUser) ToRichModel() User {
	id := a.ID
	deviceID := a.DeviceID
	createdAt := a.CreatedAt
	updatedAt := a.UpdatedAt
	deletedAt := a.DeletedAt
	return User{id, deviceID, createdAt, updatedAt, deletedAt}
}
