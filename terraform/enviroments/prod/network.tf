module "network" {
  source = "../../modules/network"

  environment = var.environment
  project     = var.project
  tags        = local.common_tags

  vpc_cidr             = "10.0.0.0/16"
  public_subnet_cidrs  = ["10.0.0.0/24", "10.0.1.0/24"]
  private_subnet_cidrs = ["10.0.10.0/24", "10.0.11.0/24"]
  az_names             = ["ap-northeast-1a", "ap-northeast-1c"]
}


