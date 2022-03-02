output "vpc" {
    value = module.vpc.vpc
}

output "public" {
    value = module.vpc.public
}

output "private" {
    value = module.vpc.private
}

output "gateway" {
    value = module.vpc.gateway
}

output "nat" {
    value = module.vpc.nat
}

output "security_groups" {
    value = module.vpc.security_groups
}

output "load_balancer" {
    value = module.vpc.alb
}

output "ecs" {
    value = module.ecs.ecs
}

output "eks" {
    value = module.eks.eks
}