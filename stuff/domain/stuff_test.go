package domain

import (
	"testing"
	"time"
)

func TestNewStuff(t *testing.T) {
	id := "id"
	name := "name"
	category := "category"
	folder := "folder"
	userID := "userID"
	fileKey := "image file Key"

	stuff := NewStuff(id, name, category, folder, userID, fileKey)

	if stuff.id != id {
		t.Fail()
	}
	if stuff.name != name {
		t.Fail()
	}
	if stuff.category != category {
		t.Fail()
	}
	if stuff.folder != folder {
		t.Fail()
	}
	if stuff.userID != userID {
		t.Fail()
	}
}

func TestToRichModel(t *testing.T) {
	id := "id"
	name := "name"
	category := "category"
	folder := "folder"
	userID := "userID"
	fileKey := "image file Key"
	createdAt := time.Now()
	updatedAt := time.Now()

	stuff := Stuff{id, name, category, folder, userID, fileKey, createdAt, updatedAt}

	anemic := AnemicStuff{id, name, category, folder, userID, fileKey, createdAt, updatedAt}
	result := anemic.ToRichModel()

	if result != stuff {
		t.Fail()
	}
}
