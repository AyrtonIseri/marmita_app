resource "aws_security_group" "webserver_backend_sg" {
  name        = "webserver-backend-sg"
  description = "Security group of the webservers backend hosting ec2"
  vpc_id      = module.vpc.vpc_id
}

resource "aws_vpc_security_group_egress_rule" "webserver_backend_allow_ssm_vpce" {
  security_group_id            = aws_security_group.webserver_backend_sg.id
  referenced_security_group_id = aws_security_group.ssm_vpce_sg.id

  from_port   = "443"
  to_port     = "443"
  ip_protocol = "tcp"
}

resource "aws_security_group" "ssm_vpce_sg" {
  name        = "ssm-vpce-sg"
  description = "Security group to be used with vpce related to the ssm service to log into the webserver"
  vpc_id      = module.vpc.vpc_id
}

resource "aws_vpc_security_group_ingress_rule" "ssm_vpce_inbound_rules" {
  security_group_id = aws_security_group.ssm_vpce_sg.id

  cidr_ipv4   = "0.0.0.0/0"
  ip_protocol = "-1"
}

resource "aws_vpc_security_group_egress_rule" "ssm_vpce_outbound_rules" {
  security_group_id = aws_security_group.ssm_vpce_sg.id

  cidr_ipv4   = "0.0.0.0/0"
  ip_protocol = "-1"
}
