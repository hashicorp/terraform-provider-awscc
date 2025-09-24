# Amazon EKS Pod Identity associations provide the ability to manage credentials for your applications, similar to the way that Amazon EC2 instance profiles provide credentials to Amazon EC2 instances.
# It associates an IAM role with a Service Account which is then associated with Pods. 
# First create IAM role for EKS Pod Identity
resource "awscc_iam_role" "pod_identity_role" {
  role_name = "eks_pod_identity_role"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Sid    = "AllowEksAuthToAssumeRoleForPodIdentity"
      Effect = "Allow"
      Principal = {
        Service = "pods.eks.amazonaws.com" # One trust policy for all EKS clusters.
      }
      Action = [
        "sts:AssumeRole",
        "sts:TagSession"
      ]
    }]
  })
  # Add policy ARNs as needed. Here is an example:  
  managed_policy_arns = ["arn:aws:iam::aws:policy/AmazonS3ReadOnlyAccess"]
}

# Associate the IAM role with a Service Account
resource "awscc_eks_pod_identity_association" "pod_identity_association_s3_readonly" {
  cluster_name    = var.cluster_name
  namespace       = var.namespace
  service_account = var.serviceaccount
  role_arn        = awscc_iam_role.pod_identity_role.arn # like: arn:aws:iam::xxxxxxxxxxxx:role/role1
}

variable "cluster_name" {
  type = string
}

variable "namespace" {
  type = string
}

variable "serviceaccount" {
  type = string
}
