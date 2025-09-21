variable "environment" {
  type        = string
  description = "デプロイ環境（例: dev, prod）"
}

variable "project" {
  type        = string
  description = "プロジェクト名（例: todo-gin）"
}

variable "tags" {
  type        = map(string)
  description = "共通タグ（Environment, Project を含めること）"
  default     = {}
}

variable "vpc_cidr" {
  type        = string
  description = "VPC の CIDR ブロック（例: 10.0.0.0/16）"
}

variable "az_names" {
  type        = list(string)
  description = "利用する AZ 名（例: ap-northeast-1a, ap-northeast-1c）"
}

variable "public_subnet_cidrs" {
  type        = list(string)
  description = "Public サブネットの CIDR 一覧（2つ想定）"
}

variable "private_subnet_cidrs" {
  type        = list(string)
  description = "Private サブネットの CIDR 一覧（2つ想定）"
}


