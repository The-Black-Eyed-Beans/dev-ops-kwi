module "vpc" {
    source = "../modules/vpc"
    vpc_cidr = var.vpc_cidr
    vpc_sg = var.vpc_sg
    public = var.public
    private = var.private
    alb = var.alb
}

module "ecs" {
    source = "../modules/ecs"
}

module "eks" {
    source = "../modules/eks"
    private = module.vpc.private
    security_groups = module.vpc.security_groups
}