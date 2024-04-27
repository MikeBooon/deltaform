package email

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"github.com/aws/aws-sdk-go/aws"
)

const (
	charset = "UTF-8"
)

func getOTPEmailTemplate(code string) *types.Message {
	htmlBody := fmt.Sprintf(`
		<h1>deltaforms</h1>
		<h2>One time passcode</h2>
		<p>%s</p>
	`, code)

	textBody := fmt.Sprintf(`deltaforms otp code: %s`, code)

	subject := "deltaform Email Verifcation Code"

	return &types.Message{
		Body: &types.Body{
			Html: &types.Content{
				Charset: aws.String(charset),
				Data:    aws.String(htmlBody),
			},
			Text: &types.Content{
				Charset: aws.String(charset),
				Data:    aws.String(textBody),
			},
		},
		Subject: &types.Content{
			Charset: aws.String(charset),
			Data:    aws.String(subject),
		},
	}
}
