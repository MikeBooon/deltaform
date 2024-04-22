package service

import (
	"database/sql"
	"time"

	"github.com/mikebooon/deltaform/internal/auth"
	"github.com/mikebooon/deltaform/internal/db/model"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

const (
	verificationExpiresMinutes = 5
)

func (s *UserService) GetNewVerificationCode(email string) (string, error) {
	code, err := auth.GenerateVerificationCode()

	if err != nil {
		return "", err
	}

	exp := time.Now().Add(verificationExpiresMinutes * time.Minute)

	var user model.User

	result := s.db.FirstOrCreate(&user, model.User{
		Email: email,
	})

	if result.Error != nil {
		return "", result.Error
	}

	codeHash, err := auth.GenerateHash(code)

	if err != nil {
		return "", err
	}

	user.VerificationCode = sql.NullString{
		String: codeHash,
		Valid:  true,
	}

	user.VerificationExpires = sql.NullTime{
		Time:  exp,
		Valid: true,
	}

	s.db.Save(&user)

	return code, nil
}
