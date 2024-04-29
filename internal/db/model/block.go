package model

type Block struct {
	ID          string `gorm:"primarykey;type:uuid;default:gen_random_uuid()"`
	Type        string
	Title       string
	Description string
	FormID      uint
	Form        Form
	Order       int
	Fields      []Field
}
