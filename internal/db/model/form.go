package model

import (
	"github.com/mikebooon/deltaform/enums"
	"gorm.io/gorm"
)

type Form struct {
	gorm.Model
	Title     string
	Block     []Block
	StatusID  enums.FormStatus
	Status    FormStatus
	UserID    string
	User      User
	Responses []Response
}
