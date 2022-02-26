variable "private" {
  type        = any
  description = "The private subnets of the VPC."
}

variable "security_groups" {
  type        = any
  description = "The security groups used in the VPC."
}