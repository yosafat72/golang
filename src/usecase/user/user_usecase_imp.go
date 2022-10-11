package user

import (
	userRepo "go-echo/src/repository/user"
	request "go-echo/src/usecase/user/request"
	response "go-echo/src/usecase/user/response"
)

type userUsecase struct {
	userRepo userRepo.UserRepo
}

func NewUserUsecase(userRepo userRepo.UserRepo) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (uc *userUsecase) FindByIdUser(in request.UserReq) (out response.UserRep, httpCode int, code string, err error) {
	data, err := uc.userRepo.FindByIdUser(in.IdUser)
	if err != nil {
		code = "ERR"
		return
	}

	out.NmUser = data.NmUser
	out.DobUser = data.DobUser
	out.GenderUser = data.GenderUser
	out.PhoneUser = data.PhoneUser
	out.EmailUser = data.EmailUser
	out.DescUser = data.DescUser

	httpCode = 200
	code = "M1"
	return
}
