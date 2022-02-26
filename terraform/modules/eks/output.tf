output "eks" {
    description = "EKS Cluster for the project."
    value = aws_eks_cluster.eks
}