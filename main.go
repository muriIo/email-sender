package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ses"
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
