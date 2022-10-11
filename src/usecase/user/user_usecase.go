package user

import (
	request "go-echo/src/usecase/user/request"
	response "go-echo/src/usecase/user/response"
)

type UserUsecase interface {
	FindByIdUser(in request.UserReq) (out response.UserRep, httpCode int, code string, err error)
	SaveUser(in request.UserSaveReq) (out response.UserRep, httpCode int, code string, err error)
}
