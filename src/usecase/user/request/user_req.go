package request

type UserReq struct {
	IdUser string `json:"idUser" binding:"required" validate:"required"`
}

type UserSaveReq struct {
	NmUser     string `json:"nmUser"`
	DobUser    string `json:"dob"`
	GenderUser string `json:"gender"`
	PhoneUser  string `json:"phone"`
	EmailUser  string `json:"email"`
	DescUser   string `json:"desc"`
}
