package migrate

import (
	"github.com/mikebooon/deltaform/domain"
	"github.com/mikebooon/deltaform/internal/db/model"
	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	db.AutoMigrate(
		&model.Form{},
		&model.User{},
		&model.InputType{},
		&model.FormStatus{},
		&model.Block{},
		&model.Field{},
		&model.Response{},
		&model.FieldResponse{},
	)

	for _, id := range domain.InputTypeOptions {
		db.FirstOrCreate(&model.InputType{ID: string(id)})
	}

	for _, id := range domain.FormStatusOptions {
		db.FirstOrCreate(&model.FormStatus{ID: string(id)})
	}
}
