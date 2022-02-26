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

output "security_groups" {
    value = module.vpc.security_groups
}