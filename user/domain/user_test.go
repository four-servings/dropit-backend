package domain

import (
	"testing"
	"time"
)

func TestNewUser(t *testing.T) {
	id := "id"
	deviceID := "deviceID"

	user := NewUser(id, deviceID)

	if user.id != id {
		t.Fail()
	}
	if user.deviceID != deviceID {
		t.Fail()
	}
}

func TestToRichModel(t *testing.T) {
	id := "id"
	deviceID := "deviceID"
	createdAt := time.Now()
	updatedAt := time.Now()

	user := User{id, deviceID, createdAt, updatedAt, nil}

	anemic := AnemicUser{id, deviceID, createdAt, updatedAt, nil}
	result := anemic.ToRichModel()

	if result != user {
		t.Fail()
	}
}
