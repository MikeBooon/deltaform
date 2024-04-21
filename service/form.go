package service

import (
	"github.com/mikebooon/deltaform/internal/db/model"
	"gorm.io/gorm"
)

type FormService struct {
	db *gorm.DB
}

func NewFormService(db *gorm.DB) *FormService {
	return &FormService{
		db: db,
	}
}

func (s *FormService) GetByID(id uint) (model.Form, error) {
	var form model.Form

	result := s.db.First(&form, id)

	if result.Error == nil {
		return form, nil
	}

	return model.Form{}, result.Error
}
