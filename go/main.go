package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/katsuokaisao/api-lambda-dynamo/dynamodb"
)

func handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	awsCfg, err := dynamodb.ProvideAWSconfig()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Internal Server Error",
		}, err
	}
	dynamodbCli, err := dynamodb.ProvideDynamoDBClient(awsCfg)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Internal Server Error",
		}, err
	}
	basicCli := dynamodb.NewBasic(dynamodbCli)
	employeeCli := dynamodb.NewEmployee(basicCli)
	employees, err := employeeCli.FindAll()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Internal Server Error",
		}, err
	}
	res, err := json.Marshal(employees)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Internal Server Error",
		}, err
	}

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(res),
	}
	return response, nil
}

func main() {
	lambda.Start(handler)
}
