package user

import (
	database "go-echo/src/adapter/database"
	models "go-echo/src/models"
	"go-echo/src/shared/utils"

	"github.com/reactivex/rxgo/v2"
)

type userRepo struct {
	mysql *database.MySql
}

func NewUserRepo(mysql *database.MySql) UserRepo {
	return &userRepo{
		mysql: mysql,
	}
}

func (s *userRepo) FindAllUsers() (out []models.User, err error) {
	result := s.mysql.Find(&out)
	err = result.Error
	return
}

func (s *userRepo) FindByIdUser(in string) (out models.User, err error) {
	err = s.mysql.Where("id_user = ?", in).Find(&out).Error
	return
}

func (s *userRepo) SaveUser(in models.User) (out models.User, err error) {
	err = s.mysql.Create(&in).Error
	return
}

func (s *userRepo) UpdateUser(in models.User) (out models.User, err error) {
	err = s.mysql.Save(&out).Error
	return
}

func (s *userRepo) DeleteUser(in string) (out models.User, err error) {
	err = s.mysql.Where("id_user = ?", in).Delete(&out).Error
	return
}

// FetchUsers implements UserRepo.
func (s *userRepo) FetchUsers() utils.ObservableResult {
	var users []models.User
	result := s.mysql.Find(&users)
	err := result.Error
	return utils.ObservableResult{
		Result: rxgo.Just(users)(),
		Error:  err,
	}
}
