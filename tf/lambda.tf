resource "aws_iam_role" "lambda" {
  name               = "lambda"
  assume_role_policy = data.aws_iam_policy_document.lambda_assume_role_policy.json
}

data "aws_iam_policy_document" "lambda_assume_role_policy" {
  statement {
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}

resource "aws_iam_role_policy_attachment" "lambda" {
  role       = aws_iam_role.lambda.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_iam_role_policy" "lambda-dynamodb" {
  name = "lambda-dynamodb"
  role = aws_iam_role.lambda.name
  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect = "Allow",
        Action = [
          "dynamodb:PartiQLSelect",
        ],
        Resource = "${aws_dynamodb_table.employees.arn}"
      }
    ]
  })
}

resource "aws_lambda_function" "lambda_go" {
  function_name = "lambda-go"
  package_type  = "Image"
  image_uri     = "${aws_ecr_repository.go_lambda.repository_url}:latest"
  role          = aws_iam_role.lambda.arn
  publish       = false
  memory_size   = 128
  timeout       = 30
}

resource "aws_lambda_permission" "apigw_lambda" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.lambda_go.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_api_gateway_rest_api.example.execution_arn}/*"
}