package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID                  string `gorm:"primarykey;type:uuid;default:gen_random_uuid()"`
	Email               string `gorm:"index,unique"`
	CreatedAt           time.Time
	VerificationCode    sql.NullString
	VerificationExpires sql.NullTime
}
