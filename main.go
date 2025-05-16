package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ses"
)

var (
	sesClient *ses.Client
)

func handleTriggering(ctx context.Context) error {
	log.Printf("This is the context: %v", ctx)

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
