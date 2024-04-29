package model

import "github.com/mikebooon/deltaform/enums"

type Field struct {
	ID          string `gorm:"primarykey;type:uuid;default:gen_random_uuid()"`
	TypeID      enums.InputType
	Type        InputType
	Title       string
	Description string
	BlockID     string
	Block       Block
	Min         int
	Max         int
	Required    bool
	Order       int
	Responses   []FieldResponse
}
