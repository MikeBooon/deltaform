package email

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
)

type EmailClient struct {
	ses    ses.Client
	sender string
}

type EmailOptions struct {
	Sender    string
	AccessKey string
	SecretKey string
}

func NewEmailClient(opts EmailOptions) *EmailClient {
	ses := ses.NewFromConfig(aws.Config{
		Credentials: credentials.NewStaticCredentialsProvider(
			opts.AccessKey,
			opts.SecretKey,
			"",
		),
		Region: "us-east-2",
	})

	return &EmailClient{
		ses:    *ses,
		sender: opts.Sender,
	}
}

func (client *EmailClient) SendOTPEmail(to string, code string) {
	template := getOTPEmailTemplate(code)

	input := &ses.SendEmailInput{
		Destination: &types.Destination{
			ToAddresses: []string{
				to,
			},
		},
		Message: template,
		Source:  aws.String(client.sender),
	}

	_, err := client.ses.SendEmail(context.TODO(), input)

	if err != nil {
		log.Fatal("Failed to send email", err.Error())
	}
}
