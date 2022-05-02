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
    value = module.vpc.micro_alb
}

output "gateway_load_balancer" {
    value = module.vpc.gate_alb
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