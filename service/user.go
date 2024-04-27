package service

import (
	"database/sql"
	"errors"
	"log"
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

func (s *UserService) VerifyVerficationCode(code string, email string) (bool, error) {
	var user model.User

	result := s.db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, ErrUserNotFoundWithEmail
		}

		return false, result.Error
	}

	if !user.VerificationCode.Valid {
		return false, nil
	}

	if time.Now().After(user.VerificationExpires.Time) {
		return false, nil
	}

	log.Println(code)

	isValid := auth.CompareHash(user.VerificationCode.String, code)

	log.Println(isValid)

	if !isValid {
		return false, nil
	}

	user.VerificationCode = sql.NullString{
		Valid: false,
	}

	user.VerificationExpires = sql.NullTime{
		Valid: false,
	}

	s.db.Save(&user)

	return true, nil
}
