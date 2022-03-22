module "vpc" {
    source = "../modules/vpc"
    vpc_cidr = var.vpc_cidr
    vpc_sg = var.vpc_sg
    public = var.public
    private = var.private
    micro_alb = var.micro_alb
    gate_alb = var.gate_alb
    route53_domain = var.route53_domain
}

module "ecs" {
    source = "../modules/ecs"
}

module "eks" {
    source = "../modules/eks"
    private = module.vpc.private
}

module "secret" {
    source = "../modules/secret"
    vpc = module.vpc.vpc
    public = module.vpc.public
    private = module.vpc.private
    app_security_group = module.vpc.microservice_sg
    gateway_security_group = module.vpc.gateway_sg
    app_alb = module.vpc.micro_alb
    gateway_alb = module.vpc.gate_alb
    ecs = module.ecs.ecs
    eks = module.eks.eks
    ecs_role = module.ecs.ecs_role
    task_role = module.ecs.task_role
    eks_role = module.eks.eks_role
}