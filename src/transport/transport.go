package transport

import (
	"go-echo/src/interface/container"
)

type Tp struct {
	UserTransport *userTransport
}

func SetupTransport(container *container.Container) *Tp {
	userTp := NewUserTransport(container.UserUsecase)

	return &Tp{
		UserTransport: userTp,
	}
}
