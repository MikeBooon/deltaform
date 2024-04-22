package service

import "gorm.io/gorm"

type ServiceRepo struct {
	FormService FormService
	UserService UserService
}

func NewServiceRepo(db gorm.DB) *ServiceRepo {
	return &ServiceRepo{
		FormService: FormService{
			db: &db,
		},
		UserService: UserService{
			db: &db,
		},
	}
}
