package command

import (
	"errors"
	"github/four-servings/dropit-backend/stuff/infra"
)

type (
	// Bus command bus interface
	Bus interface {
		Handle(command interface{})
	}

	busImplement struct {
		createStuffHandler createStuffHandler
	}
)

// NewBus create bus instance
func NewBus(repository infra.StuffRepository) Bus {
	createStuffHandler := newCreateStuffHandler(repository)
	return &busImplement{createStuffHandler}
}

// Handle handle given command
func (b *busImplement) Handle(givenCommand interface{}) {
	switch givenCommand := givenCommand.(type) {
	case *CreateStuff:
		b.createStuffHandler.handle(givenCommand)
		return
	default:
		panic(errors.New("invalid command"))
	}
}
