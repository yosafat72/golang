package user

import (
	"go-echo/src/usecase/user/request"
	"net/http"

	resp "go-echo/src/shared/utils"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func FindByIdUserValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req request.UserReq

		err := c.Bind(&req)

		if err != nil {
			return resp.SetResponse(c, http.StatusBadRequest, "Error validation", nil)
		}

		v := validator.New()
		if err := v.Struct(req); err != nil {
			return resp.SetResponse(c, http.StatusBadRequest, "Error validation", nil)
		}
		c.Set("body", req)
		return next(c)
	}
}
