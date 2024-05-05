package dynamodb

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func ProvideAWSconfig() (aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion("ap-northeast-1"),
	)
	if err != nil {
		return aws.Config{}, fmt.Errorf("failed to load configuration, %v", err)
	}
	return cfg, nil
}

func ProvideDynamoDBClient(cfg aws.Config) (*dynamodb.Client, error) {
	client := dynamodb.NewFromConfig(cfg)
	return client, nil
}
