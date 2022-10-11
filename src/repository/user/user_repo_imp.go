package user

import (
	database "go-echo/src/adapter/database"
	models "go-echo/src/models"
)

type userRepo struct {
	postgres *database.PostgreSQL
}

func NewUserRepo(postgres *database.PostgreSQL) UserRepo {
	return &userRepo{
		postgres: postgres,
	}
}

func (s *userRepo) FindByIdUser(in string) (out models.User, err error) {
	err = s.postgres.Where("id_user = ?", in).Find(&out).Error
	err = s.postgres.Error
	return
}