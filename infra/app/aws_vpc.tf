module "vpc" {
  source = "terraform-aws-modules/vpc/aws"
  version = ">= 5.9.0"

  name = local.vpc_name
  cidr = local.vpc_cidr

  azs             = local.vpc_azs
  private_subnets = local.private_subnets
  public_subnets  = local.public_subnets

  private_subnet_names = [for index, subnet in local.private_subnets : "marmita-private-subnet-${index}"]
  public_subnet_names  = [for index, subnet in local.public_subnets : "marmita-public-subnet-${index}"]

  single_nat_gateway = true
  enable_nat_gateway = true
  enable_vpn_gateway = false

}
