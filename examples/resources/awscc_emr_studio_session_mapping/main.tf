data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# IAM role for EMR Studio
resource "aws_iam_role" "emr_studio_role" {
  name = "emr-studio-basic-role"
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "elasticmapreduce.amazonaws.com"
        }
      }
    ]
  })
  tags = {
    "Modified By" = "AWSCC"
  }
}

# Basic policy for EMR Studio
data "aws_iam_policy_document" "emr_studio_basic" {
  statement {
    effect = "Allow"
    actions = [
      "emr-containers:ListVirtualClusters",
      "emr-containers:DescribeVirtualCluster",
      "elasticmapreduce:ListInstances",
      "elasticmapreduce:DescribeCluster",
      "elasticmapreduce:ListSteps",
      "elasticmapreduce:ListClusters"
    ]
    resources = ["*"]
  }
}

resource "aws_iam_policy" "emr_studio_basic" {
  name   = "emr-studio-basic-policy"
  policy = data.aws_iam_policy_document.emr_studio_basic.json
  tags = {
    "Modified By" = "AWSCC"
  }
}

resource "aws_iam_role_policy_attachment" "emr_studio_basic" {
  policy_arn = aws_iam_policy.emr_studio_basic.arn
  role       = aws_iam_role.emr_studio_role.name
}

# Session policy for EMR Studio
data "aws_iam_policy_document" "session_policy" {
  statement {
    effect = "Allow"
    actions = [
      "emr-containers:ListVirtualClusters",
      "emr-containers:DescribeVirtualCluster",
      "elasticmapreduce:ListInstances",
      "elasticmapreduce:DescribeCluster",
      "elasticmapreduce:ListSteps"
    ]
    resources = ["*"]
  }
}

resource "aws_iam_policy" "session_policy" {
  name   = "emr-studio-session-policy"
  policy = data.aws_iam_policy_document.session_policy.json
  tags = {
    "Modified By" = "AWSCC"
  }
}

# EMR Studio
resource "awscc_emr_studio" "example" {
  auth_mode                   = "IAM"
  default_s3_location         = "s3://emr-studio-${data.aws_caller_identity.current.account_id}-${data.aws_region.current.name}/example"
  engine_security_group_id    = "sg-0123456789abcdef0" # Replace with your security group ID
  name                        = "example-studio"
  service_role                = aws_iam_role.emr_studio_role.arn
  subnet_ids                  = ["subnet-0123456789abcdef0"] # Replace with your subnet ID
  vpc_id                      = "vpc-0123456789abcdef0"      # Replace with your VPC ID
  workspace_security_group_id = "sg-0123456789abcdef0"       # Replace with your security group ID
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# EMR Studio Session Mapping
resource "awscc_emr_studio_session_mapping" "example" {
  identity_name      = "example-user"
  identity_type      = "USER"
  session_policy_arn = aws_iam_policy.session_policy.arn
  studio_id          = awscc_emr_studio.example.id
}