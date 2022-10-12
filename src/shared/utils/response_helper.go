package utils

import (
	echo "github.com/labstack/echo/v4"
)

type JSONResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SetResponse(c echo.Context, statusCode int, msg string, data interface{}) error {
	return c.JSON(statusCode, JSONResponse{
		Code:    statusCode,
		Message: msg,
		Data:    data,
	})
}
