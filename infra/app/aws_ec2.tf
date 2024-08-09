resource "aws_instance" "marmita_backend_webserver" {
  ami                         = "ami-0ae8f15ae66fe8cda"
  associate_public_ip_address = false

  iam_instance_profile = aws_iam_instance_profile.ssm_managed_ec2_instance_profile_core.name
  instance_type        = "t2.micro"
  security_groups      = [aws_security_group.webserver_backend_sg.id]
  source_dest_check    = false
  subnet_id            = module.vpc.private_subnets[0]
  tenancy              = "default"

  tags = {
    Name = "marmita-backend-webserver"
  }
}

resource "aws_iam_instance_profile" "ssm_managed_ec2_instance_profile_core" {
  name = "SSMManagedEC2InstanceCore"
  role = aws_iam_role.ec2_ssm_permission_role.name
}
