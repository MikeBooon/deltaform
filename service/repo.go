package service

import "gorm.io/gorm"

type ServiceRepo struct {
	FormService FormService
}

func NewServiceRepo(db gorm.DB) *ServiceRepo {
	return &ServiceRepo{
		FormService: FormService{
			db: &db,
		},
	}
}
