output "vpc_id" {
  value       = module.network.vpc_id
  description = "VPC ID"
}

output "public_subnet_ids" {
  value       = module.network.public_subnet_ids
  description = "Public サブネット ID"
}

output "private_subnet_ids" {
  value       = module.network.private_subnet_ids
  description = "Private サブネット ID"
}

output "nat_gateway_id" {
  value       = module.network.nat_gateway_id
  description = "NAT Gateway ID"
}


