
output "remote_state_bucket_id" {
  value = aws_s3_bucket.remote_state_bucket.id
}

output "remote_state_dynamo_table_id" {
  value = aws_dynamodb_table.tf_state_dynamo_db_table.id
}

output "remote_state_kms_key_id" {
  value = aws_kms_key.tf_state_kms_key.id
}
