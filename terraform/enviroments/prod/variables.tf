variable "aws_region" {
  type        = string
  description = "AWSリージョン（例: ap-northeast-1）"
  default     = "ap-northeast-1"
}

variable "environment" {
  type        = string
  description = "環境名（例: prod, dev）"
  default     = "prod"
}

variable "project" {
  type        = string
  description = "プロジェクト名"
  default     = "todo-gin"
}


