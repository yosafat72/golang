package container

import (
	"go-echo/src/adapter/database"
	userRepo "go-echo/src/repository/user"
	userUsecase "go-echo/src/usecase/user"
)

type Container struct {
	UserUsecase userUsecase.UserUsecase
}

func SetupContainer() Container {

	//Setup database
	mysql := database.SetupMysql()

	//init repo
	userRepository := userRepo.NewUserRepo(mysql)

	//init usecase
	userCase := userUsecase.NewUserUsecase(userRepository)

	return Container{
		UserUsecase: userCase,
	}
}
