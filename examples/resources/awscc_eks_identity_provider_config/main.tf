# Create VPC for EKS
resource "aws_vpc" "eks" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true

  tags = {
    Name = "eks-vpc"
  }
}

# Create public subnets
resource "aws_subnet" "public_1" {
  vpc_id            = aws_vpc.eks.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "us-west-2a"

  tags = {
    Name = "eks-public-1"
  }
}

resource "aws_subnet" "public_2" {
  vpc_id            = aws_vpc.eks.id
  cidr_block        = "10.0.2.0/24"
  availability_zone = "us-west-2b"

  tags = {
    Name = "eks-public-2"
  }
}

# IAM role for EKS cluster
resource "aws_iam_role" "eks_cluster" {
  name = "eks-cluster-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Principal = {
          Service = "eks.amazonaws.com"
        }
        Action = "sts:AssumeRole"
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "eks_cluster_policy" {
  policy_arn = "arn:aws:iam::aws:policy/AmazonEKSClusterPolicy"
  role       = aws_iam_role.eks_cluster.name
}

# EKS Cluster
resource "aws_eks_cluster" "demo" {
  name     = "demo-cluster"
  role_arn = aws_iam_role.eks_cluster.arn

  vpc_config {
    subnet_ids = [
      aws_subnet.public_1.id,
      aws_subnet.public_2.id
    ]
  }

  depends_on = [
    aws_iam_role_policy_attachment.eks_cluster_policy
  ]
}

# EKS Identity Provider Config using AWSCC
resource "awscc_eks_identity_provider_config" "example" {
  cluster_name = aws_eks_cluster.demo.name
  type         = "oidc"

  identity_provider_config_name = "example-oidc-config"

  oidc = {
    client_id      = "example-client"
    issuer_url     = "https://example.com"
    groups_claim   = "groups"
    groups_prefix  = "example-group-"
    username_claim = "email"
    required_claims = [{
      key   = "email_verified"
      value = "true"
    }]
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}