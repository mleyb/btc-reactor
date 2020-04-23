package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"log"
)

// PublishSMS sends an SMS message
func PublishSMS(message string, phoneNumber string) {
	sess := session.Must(session.NewSession())

	svc := sns.New(sess)

	params := &sns.PublishInput{
		Message:     aws.String(message),
		PhoneNumber: aws.String(phoneNumber),
	}

	resp, err := svc.Publish(params)

	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println(resp)
}
