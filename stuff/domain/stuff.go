package domain

import (
	"time"
)

type (
	// Stuff stuff model
	Stuff struct {
		id        string
		name      string
		category  string
		createdAt time.Time
		updatedAt time.Time
	}

	// AnemicStuff anemic stuff model
	AnemicStuff struct {
		ID        string
		Name      string
		Category  string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)

// ID get stuff id
func (stuff *Stuff) ID() string {
	return stuff.id
}

// Name get stuff name
func (stuff *Stuff) Name() string {
	return stuff.name
}

// Category get stuff category
func (stuff *Stuff) Category() string {
	return stuff.category
}

// CreatedAt get stuff createdAt
func (stuff *Stuff) CreatedAt() time.Time {
	return stuff.createdAt
}

// UpdatedAt get stuff updatedAt
func (stuff *Stuff) UpdatedAt() time.Time {
	return stuff.updatedAt
}

// NewStuff create new Stuff
func NewStuff(id, name, category string) Stuff {
	return Stuff{id, name, category, time.Now(), time.Now()}
}

// ToRichModel create rich stuff model from anemic model
func (a *AnemicStuff) ToRichModel() Stuff {
	id := a.ID
	name := a.Name
	category := a.Category
	createdAt := a.CreatedAt
	updatedAt := a.UpdatedAt
	return Stuff{id, name, category, createdAt, updatedAt}
}