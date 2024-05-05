resource "aws_api_gateway_account" "example" {
  cloudwatch_role_arn = aws_iam_role.example_apigateway_role.arn
}

resource "aws_api_gateway_deployment" "example_test" {
  rest_api_id = aws_api_gateway_rest_api.example.id
}

resource "aws_api_gateway_stage" "example_test" {
  rest_api_id   = aws_api_gateway_rest_api.example.id
  deployment_id = aws_api_gateway_deployment.example_test.id
  stage_name    = "test"
}

resource "aws_iam_role" "example_apigateway_role" {
  name = "example-apigateway"
  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect = "Allow",
        Principal = {
          Service = "apigateway.amazonaws.com"
        },
        Action = "sts:AssumeRole"
      }
    ]
  })
}

resource "aws_iam_role_policy" "example_apigateway_role_policy" {
  name = "example-apigateway"
  role = aws_iam_role.example_apigateway_role.id
  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect = "Allow",
        Action = [
          "logs:CreateLogGroup",
          "logs:CreateLogStream",
          "logs:DescribeLogGroups",
          "logs:DescribeLogStreams",
          "logs:PutLogEvents",
          "logs:GetLogEvents",
          "logs:FilterLogEvents"
        ],
        Resource = "*"
      }
    ]
  })
}

resource "aws_api_gateway_rest_api" "example" {
  name = "example"
}

resource "aws_api_gateway_resource" "employee" {
  rest_api_id = aws_api_gateway_rest_api.example.id
  parent_id   = aws_api_gateway_rest_api.example.root_resource_id
  path_part   = "employee"
}

resource "aws_api_gateway_method" "employee_get" {
  rest_api_id   = aws_api_gateway_rest_api.example.id
  resource_id   = aws_api_gateway_resource.employee.id
  http_method   = "GET"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "employee_get" {
  rest_api_id             = aws_api_gateway_rest_api.example.id
  resource_id             = aws_api_gateway_resource.employee.id
  http_method             = aws_api_gateway_method.employee_get.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.lambda_go.invoke_arn
}
