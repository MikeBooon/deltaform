package model

type FieldResponse struct {
	ID         string `gorm:"primarykey;type:uuid;default:gen_random_uuid()"`
	ResponseID string
	Response   Response
	FieldID    string
	Field      Field
	Value      string
}
