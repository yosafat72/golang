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

func (s *userRepo) FindAllUsers() (out []models.User, err error) {
	result := s.postgres.Find(&out)
	err = result.Error
	return
}

func (s *userRepo) FindByIdUser(in string) (out models.User, err error) {
	err = s.postgres.Where("id_user = ?", in).Find(&out).Error
	return
}

func (s *userRepo) SaveUser(in models.User) (out models.User, err error) {
	err = s.postgres.Create(&in).Error
	return
}

func (s *userRepo) UpdateUser(in models.User) (out models.User, err error) {
	err = s.postgres.Save(&out).Error
	return
}

func (s *userRepo) DeleteUser(in string) (out models.User, err error) {
	err = s.postgres.Where("id_user = ?", in).Delete(&out).Error
	return
}
