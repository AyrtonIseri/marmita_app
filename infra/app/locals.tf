locals {
    vpc_cidr = "10.0.0.0/16"
    vpc_name = "marmita-vpc"
    num_azs = 2
    vpc_azs = slice(data.aws_availability_zones.all_azs.names, 0, local.num_azs)

    private_subnets = [for index, az in local.vpc_azs : cidrsubnet(local.vpc_cidr, 8, index)]
    public_subnets = [for index, az in local.vpc_azs : cidrsubnet(local.vpc_cidr, 8, index + local.num_azs)]
}
