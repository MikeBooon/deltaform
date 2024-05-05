package model

import (
	"github.com/mikebooon/deltaform/domain"
	"gorm.io/gorm"
)

type Form struct {
	gorm.Model
	Title     string
	Block     []Block
	StatusID  domain.FormStatus
	Status    FormStatus
	UserID    string
	User      User
	Responses []Response
}
