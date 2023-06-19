package user

import (
	models "go-echo/src/models"
	"go-echo/src/shared/utils"
)

type UserRepo interface {
	FindAllUsers() (out []models.User, err error)
	FindByIdUser(in string) (out models.User, err error)
	SaveUser(in models.User) (out models.User, err error)
	UpdateUser(in models.User) (out models.User, err error)
	DeleteUser(in string) (out models.User, err error)

	//With RxGO
	FetchUsers() utils.ObservableResult
}
