resource "aws_security_group" "webserver_backend_sg" {
  name        = "webserver-backend-sg"
  description = "Security group of the webservers backend hosting ec2"
  vpc_id      = module.vpc.vpc_id
}

