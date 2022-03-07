variable "vpc_cidr" {
  type        = string
  description = "CIDR block for the VPC."
}

variable "vpc_sg" {
  type        = any
  description = "Security groups used within the VPC."
}

variable "public" {
  type        = map(object({
    cidr_block = string
    tag_name = string
    az = string
  }))
  description = "Subnets used for the VPC."
}

variable "private" {
  type        = map(object({
    cidr_block = string
    tag_name = string
    az = string
  }))
  description = "Subnets used for the VPC."
}

variable "alb" {
  type = string
  description = "Name for the Application Load Balancer."
}
