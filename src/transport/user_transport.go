package transport

import (
	userUsecase "go-echo/src/usecase/user"

	resp "go-echo/src/shared/utils"
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

		ctx.Bind(&userId)

		res, httpStatusCode, code, err := tp.UserUC.FindByIdUser(userId)
		if err != nil {
			return resp.SetResponse(ctx, httpStatusCode, code, nil)
		}

		return resp.SetResponse(ctx, 200, code, res)
	}
}
