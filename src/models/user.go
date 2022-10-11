package models

import (
	constants "go-echo/src/shared/constant"
	"time"
)

type User struct {
	Id         int64     `gorm:"primary_key;column:id"`
	IdUser     int64     `gorm:"column:id_user"`
	NmUser     string    `gorm:"column:nm_user"`
	DobUser    string    `gorm:"column:dob_user"`
	GenderUser string    `gorm:"column:gender_user"`
	PhoneUser  string    `gorm:"column:phone_user"`
	EmailUser  string    `gorm:"column:email_user"`
	DescUser   string    `gorm:"column:desc_user"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}

func (User) TableName() string {
	return constants.TBLUSER
}
