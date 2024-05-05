package dynamodb

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/katsuokaisao/api-lambda-dynamo/domain"
	"github.com/katsuokaisao/api-lambda-dynamo/repository"
)

type employee struct {
	cli repository.DynamoDBBasic
}

func NewEmployee(cli repository.DynamoDBBasic) repository.Employee {
	return &employee{
		cli: cli,
	}
}

func (e *employee) FindAll() ([]domain.Employee, error) {
	pageSize := int32(10)
	query := "SELECT * FROM employees"
	responseItems, err := e.cli.FindAll(query, pageSize)
	if err != nil {
		return nil, fmt.Errorf("failed to find all, %v", err)
	}
	employees := make([]domain.Employee, 0, len(responseItems))
	if err = attributevalue.UnmarshalListOfMaps(responseItems, &employees); err != nil {
		return nil, fmt.Errorf("failed to unmarshal list of maps, %v", err)
	}
	return employees, nil
}
