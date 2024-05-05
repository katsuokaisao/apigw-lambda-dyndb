resource "aws_dynamodb_table" "employees" {
  name           = "employees"
  hash_key       = "id"
  billing_mode   = "PROVISIONED"
  read_capacity  = 1
  write_capacity = 1
  attribute {
    name = "id"
    type = "S"
  }
}

resource "aws_dynamodb_table_item" "employee" {
  table_name = aws_dynamodb_table.employees.name
  hash_key   = aws_dynamodb_table.employees.hash_key
  item = jsonencode({
    id         = { S = "1" }
    first_name = { S = "Taro" }
    last_name  = { S = "Momo" },
    office     = { S = "Nagoya" }
  })
}
