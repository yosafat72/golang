package user

import (
	models "go-echo/src/models"
)

type UserRepo interface {
	FindByIdUser(in string) (out models.User, err error)
}