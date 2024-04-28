package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	tokenExpiry = 24 * time.Hour
)

var (
	secret = os.Getenv("JWT_SECRET")
)

type JWTBody struct {
	Email string
	ID    string
}

func GenerateJWT(body JWTBody) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": body.Email,
		"uid":   body.ID,
		"iss":   "deltaform",
		"aud":   "deltaform",
		"nbf":   time.Now().Unix(),
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(tokenExpiry).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string) (JWTBody, error) {
	var body JWTBody

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil
	})

	if err != nil {
		return body, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		body.Email = claims["email"].(string)
		body.ID = claims["uid"].(string)
	} else {
		return body, errors.New("invalid token")
	}

	return body, nil
}
