package user

import (
	"go-echo/src/models"
	userRepo "go-echo/src/repository/user"
	request "go-echo/src/usecase/user/request"
	response "go-echo/src/usecase/user/response"
	"math/rand"
)

type userUsecase struct {
	userRepo userRepo.UserRepo
}

func NewUserUsecase(userRepo userRepo.UserRepo) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (uc *userUsecase) FindAllUsers() (out []response.UserRep, httpCode int, code string, err error) {
	data, err := uc.userRepo.FindAllUsers()
	if err != nil {
		code = "ERR"
		return
	}

	for _, s := range data {
		resp := response.UserRep{NmUser: s.NmUser, DobUser: s.DobUser, GenderUser: s.GenderUser, PhoneUser: s.PhoneUser, EmailUser: s.EmailUser, DescUser: s.DescUser}
		out = append(out, resp)
	}

	httpCode = 200
	code = "M1"
	return
}

func (uc *userUsecase) FindByIdUser(in request.UserReq) (out response.UserRep, httpCode int, code string, err error) {
	data, err := uc.userRepo.FindByIdUser(in.IdUser)
	if err != nil {
		code = "ERR"
		return
	}

	resp := response.UserRep{NmUser: data.NmUser, DobUser: data.DobUser, GenderUser: data.GenderUser, PhoneUser: data.PhoneUser, EmailUser: data.EmailUser, DescUser: data.DescUser}
	out = resp

	httpCode = 200
	code = "M1"
	return
}

func (uc *userUsecase) SaveUser(in request.UserSaveReq) (out response.UserRep, httpCode int, code string, err error) {
	var data models.User

	data.IdUser = int64(rand.Intn(100))
	data.NmUser = in.NmUser
	data.DobUser = in.DobUser
	data.GenderUser = in.GenderUser
	data.PhoneUser = in.PhoneUser
	data.EmailUser = in.EmailUser
	data.DescUser = in.DescUser

	_, err = uc.userRepo.SaveUser(data)
	if err != nil {
		code = "ERR"
		return
	}

	code = "Berhasil menyimpan data!"

	return
}
