package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
)

var (
	sesClient *ses.Client
)

func handleTriggering(ctx context.Context, s3Event events.S3Event) error {
	for _, record := range s3Event.Records {
		s3 := record.S3
		log.Printf(
			"S3 event received: bucket=%s, key=%s, size=%d bytes",
			s3.Bucket.Name,
			s3.Object.Key,
			s3.Object.Size,
		)

		subject := fmt.Sprintf("Object received in the bucket: %s", s3.Bucket.Name)
		message := fmt.Sprintf("The file name is: %s and the size is %d", s3.Object.Key, s3.Object.Size)

		input := ses.SendEmailInput{
			Source: aws.String("murilomecr@outlook.com"),
			Destination: &types.Destination{
				ToAddresses: []string{"murilomecr@gmail.com"},
			},
			Message: &types.Message{
				Subject: &types.Content{
					Data: aws.String(subject),
				},
				Body: &types.Body{
					Text: &types.Content{
						Data: aws.String(message),
					},
				},
			},
		}

		result, err := sesClient.SendEmail(ctx, &input)

		if err != nil {
			log.Fatalf("Failed to send email: %v", err)
		}

		fmt.Println("Email sent successfully, Message ID:", *result.MessageId)
	}
	return nil
}

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	sesClient = ses.NewFromConfig(cfg)
}

func main() {
	lambda.Start(handleTriggering)
}
