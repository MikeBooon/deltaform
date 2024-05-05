package migrate

import (
	"log"

	"github.com/mikebooon/deltaform/domain"
	"github.com/mikebooon/deltaform/internal/db/model"
	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.Form{},
		&model.User{},
		&model.InputType{},
		&model.FormStatus{},
		&model.Block{},
		&model.Field{},
		&model.Response{},
		&model.FieldResponse{},
	)

	if err != nil {
		log.Fatal("Failed to migrate database")
		panic(err)
	}

	for _, id := range domain.InputTypeOptions {
		db.FirstOrCreate(&model.InputType{ID: string(id)})
	}

	for _, id := range domain.FormStatusOptions {
		db.FirstOrCreate(&model.FormStatus{ID: string(id)})
	}
}
