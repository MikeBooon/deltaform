package email

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
)

const (
	charset = "UTF-8"
)

func getOTPEmailTemplate(code string) *ses.Message {
	htmlBody := fmt.Sprintf(`
		<h1>deltaforms</h1>
		<h2>One time passcode</h2>
		<p>%s</p>
	`, code)

	textBody := fmt.Sprintf(`deltaforms otp code: %s`, code)

	subject := "deltaform Email Verifcation Code"

	return &ses.Message{
		Body: &ses.Body{
			Html: &ses.Content{
				Charset: aws.String(charset),
				Data:    aws.String(htmlBody),
			},
			Text: &ses.Content{
				Charset: aws.String(charset),
				Data:    aws.String(textBody),
			},
		},
		Subject: &ses.Content{
			Charset: aws.String(charset),
			Data:    aws.String(subject),
		},
	}
}
