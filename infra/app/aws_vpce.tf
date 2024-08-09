resource "aws_vpc_endpoint" "ssm_vpce" {
    vpc_id = module.vpc.vpc_id
    service_name = "com.amazonaws.${module.env.aws_provider.region}.ssm"

    vpc_endpoint_type = "Interface"

    security_group_ids = [aws_security_group.ssm_vpce_sg.id]
    private_dns_enabled = true
    subnet_ids = [ module.vpc.private_subnets[0] ]
}

resource "aws_vpc_endpoint" "ssm_message_vpce" {
    vpc_id = module.vpc.vpc_id
    service_name = "com.amazonaws.${module.env.aws_provider.region}.ssmmessages"

    vpc_endpoint_type = "Interface"

    security_group_ids = [aws_security_group.ssm_vpce_sg.id]
    private_dns_enabled = true
    subnet_ids = [ module.vpc.private_subnets[0] ]
}

