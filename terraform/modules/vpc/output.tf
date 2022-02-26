output "vpc" {
    description = "VPC for project."
    value = aws_vpc.vpc
}

output "public" {
    description = "Public subnets of the VPC."
    value = aws_subnet.public
}

output "private" {
    description = "Private subnets of the VPC."
    value = aws_subnet.private
}

output "gateway" {
    description = "Internet gateway of the VPC."
    value = aws_internet_gateway.gateway
}

output "security_groups" {
    description = "Security groups for the VPC."
    value = aws_security_group.sg
}

output "alb" {
    description = "Application Load Balancer for the VPC."
    value = aws_lb.alb
}