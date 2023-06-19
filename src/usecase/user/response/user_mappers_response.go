package response

import "go-echo/src/models"

type UserResponseMapper struct{}

func NewUserResponseMapper() *UserResponseMapper {
	return &UserResponseMapper{}
}

func (m *UserResponseMapper) ToResponse(user models.User) *UserRep {
	return &UserRep{
		NmUser:     user.NmUser,
		DobUser:    user.DobUser,
		GenderUser: user.GenderUser,
		PhoneUser:  user.PhoneUser,
		EmailUser:  user.EmailUser,
		DescUser:   user.DescUser,
	}
}
