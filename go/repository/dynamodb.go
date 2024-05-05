package repository

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/katsuokaisao/api-lambda-dynamo/domain"
)

type DynamoDBBasic interface {
	ExecuteStatement(statement string, params []types.AttributeValue) (*dynamodb.ExecuteStatementOutput, error)
	BatchExecuteStatement(statements []types.BatchStatementRequest) ([]types.BatchStatementResponse, error)
	FindAll(query string, pageSize int32) ([]map[string]types.AttributeValue, error)
}

type Employee interface {
	FindAll() ([]domain.Employee, error)
}
