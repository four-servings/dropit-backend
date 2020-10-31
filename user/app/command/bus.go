package command

import (
	"errors"
	"github/four-servings/dropit-backend/user/infra"
)

type (
	// Bus command bus interface
	Bus interface {
		Handle(command interface{})
	}

	// busImplement command bus implement
	busImplement struct {
		createUserHandler createUserHandler
	}
)

// NewBus create bus instance
func NewBus(repository infra.UserRepository) Bus {
	createUserHandler := newCreateUserHandler(repository)
	return &busImplement{createUserHandler}
}

// Handle handle given command
func (b *busImplement) Handle(givenCommand interface{}) {
	switch givenCommand := givenCommand.(type) {
	case *CreateUser:
		b.createUserHandler.handle(givenCommand)
		return
	default:
		panic(errors.New("invalid command"))
	}
}
