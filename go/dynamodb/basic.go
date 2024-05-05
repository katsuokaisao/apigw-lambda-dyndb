package dynamodb

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/katsuokaisao/api-lambda-dynamo/repository"
)

type basic struct {
	cli *dynamodb.Client
}

func NewBasic(cli *dynamodb.Client) repository.DynamoDBBasic {
	return &basic{
		cli: cli,
	}
}

func (b *basic) ExecuteStatement(statement string, params []types.AttributeValue) (*dynamodb.ExecuteStatementOutput, error) {
	input := &dynamodb.ExecuteStatementInput{
		Statement: aws.String(statement),
	}
	if len(params) > 0 {
		input.Parameters = params
	}

	response, err := b.cli.ExecuteStatement(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("failed to execute statement, %v", err)
	}
	return response, err
}

func (b *basic) FindAll(query string, pageSize int32) ([]map[string]types.AttributeValue, error) {
	res := make([]map[string]types.AttributeValue, 0, 1024)

	var nextToken *string
	for {
		input := &dynamodb.ExecuteStatementInput{
			Statement: aws.String(query),
			Limit:     aws.Int32(pageSize),
			NextToken: nextToken,
		}
		response, err := b.cli.ExecuteStatement(context.TODO(), input)
		if err != nil {
			return nil, fmt.Errorf("failed to execute statement, %v", err)
		}
		if len(response.Items) > 0 {
			res = append(res, response.Items...)
		}
		if response.NextToken == nil {
			break
		}
	}

	return res, nil
}

func (b *basic) BatchExecuteStatement(statements []types.BatchStatementRequest) ([]types.BatchStatementResponse, error) {
	input := &dynamodb.BatchExecuteStatementInput{
		Statements: statements,
	}

	response, err := b.cli.BatchExecuteStatement(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("failed to execute statement, %v", err)
	}
	return response.Responses, err
}
