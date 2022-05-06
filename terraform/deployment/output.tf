output "vpc" {
    value = module.vpc.vpc.tags.Name
}

output "vpc_cidr" {
    value = module.vpc.vpc.cidr_block
}

output "public" {
    value = module.vpc.public
}

output "private" {
    value = module.vpc.private
}

output "app_security_group" {
    value = module.vpc.microservice_sg
}

output "gateway_security_group" {
    value = module.vpc.gateway_sg
}

output "kubernetes_security_group" {
    value = module.vpc.kubernetes_sg
}

output "app_load_balancer" {
    value = module.vpc.micro_alb.name
}

output "gateway_load_balancer" {
    value = module.vpc.gate_alb.name
}

output "ecs" {
    value = module.ecs.ecs
}

output "eks" {
    value = module.eks.eks
}

output "resources_secrets" {
    value = module.secret.resources
}