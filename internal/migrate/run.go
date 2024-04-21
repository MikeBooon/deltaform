package migrate

import (
	"github.com/mikebooon/deltaform/internal/db/model"
	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	db.AutoMigrate(
		&model.Form{},
	)
}
