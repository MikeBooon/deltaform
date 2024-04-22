package auth

import (
	"crypto/rand"
	"math/big"
)

const (
	verficationCharset     = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	verificationCodeLength = 10
)

func GenerateVerificationCode() (string, error) {
	code := make([]byte, verificationCodeLength)
	for i := 0; i < verificationCodeLength; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(verficationCharset))))
		if err != nil {
			return "", err
		}
		code[i] = verficationCharset[num.Int64()]
	}

	return string(code), nil
}
