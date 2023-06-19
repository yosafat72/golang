package transport

import (
	userUsecase "go-echo/src/usecase/user"
	"net/http"

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

func (tp *userTransport) FindAllUsers() func(echo.Context) error {
	return func(ctx echo.Context) error {
		// body := ctx.Get("body").(userReq.UserReq)

		res, httpStatusCode, code, err := tp.UserUC.FindAllUsers()
		if err != nil {
			return resp.SetResponse(ctx, httpStatusCode, code, nil)
		}

		return resp.SetResponse(ctx, 200, code, res)
	}
}

func (tp *userTransport) FindByIdUser() func(echo.Context) error {
	return func(ctx echo.Context) error {
		body := ctx.Get("body").(userReq.UserReq)

		res, httpStatusCode, code, err := tp.UserUC.FindByIdUser(body)
		if err != nil {
			return resp.SetResponse(ctx, httpStatusCode, code, nil)
		}

		return resp.SetResponse(ctx, 200, code, res)
	}
}

func (tp *userTransport) SaveUser() func(echo.Context) error {
	return func(ctx echo.Context) error {
		var user userReq.UserSaveReq

		ctx.Bind(&user)

		_, httpStatusCode, code, err := tp.UserUC.SaveUser(user)
		if err != nil {
			return resp.SetResponse(ctx, httpStatusCode, code, nil)
		}

		return resp.SetResponse(ctx, 200, code, nil)
	}
}

func (tp *userTransport) FetchUsers(c echo.Context) error {

	observable := tp.UserUC.FetchUsers()

	ch := observable.Result.Observe()
	item := <-ch

	if item.E != nil {
		return resp.SetResponse(c, http.StatusInternalServerError, observable.Error.Error(), nil)
	}

	return resp.SetResponse(c, http.StatusOK, "Success", item.V)

}
