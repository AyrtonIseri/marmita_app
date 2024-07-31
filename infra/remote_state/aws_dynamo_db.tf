resource "aws_dynamodb_table" "tf_state_dynamo_db_table" {
  hash_key       = "LockID"
  name           = "marmita_app_remote_state"
  read_capacity  = 1
  write_capacity = 1

  attribute {
    name = "LockID"
    type = "S"
  }
}
