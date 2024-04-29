package model

import "time"

type Response struct {
	ID        string `gorm:"primarykey;type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time
	FormID    string
	Form      Form
}
