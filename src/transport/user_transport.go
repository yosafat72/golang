package transport

import (
	"fmt"
	userUsecase "go-echo/src/usecase/user"

	userReq "go-echo/src/usecase/user/request"

	"github.com/labstack/echo/v4"
)

type userTransport struct {
	UserUC userUsecase.UserUsecase
}

func NewUserTransport(userUsecase userUsecase.UserUsecase) *userTransport {
	return &userTransport{
		UserUC: userUsecase,
	}
}

func (tp *userTransport) FindByIdUser() func(echo.Context) error {
	return func(ctx echo.Context) error {
		var userId userReq.UserReq

		name := ctx.FormValue("idUser")

		fmt.Println(name)

		res, httpStatusCode, code, err := tp.UserUC.FindByIdUser(userId)
		_ = code
		if err != nil {
			return ctx.JSON(httpStatusCode, res)
		}

		return ctx.JSON(200, res)
	}
}
