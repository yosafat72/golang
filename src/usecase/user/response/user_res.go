package response

type UserRep struct {
	NmUser     string `json:"NmUser"`
	DobUser    string `json:"DobUser"`
	GenderUser string `json:"GenderUser"`
	PhoneUser  string `json:"PhoneUser"`
	EmailUser  string `json:"EmailUser"`
	DescUser   string `json:"DescUser"`
}
