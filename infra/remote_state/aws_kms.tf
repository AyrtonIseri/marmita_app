resource "aws_kms_key" "tf_state_kms_key" {
    description = "This key will be used to encrypt the remote state of this account's terraform"
    deletion_window_in_days = 7
}
