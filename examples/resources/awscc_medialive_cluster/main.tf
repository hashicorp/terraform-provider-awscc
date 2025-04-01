# Create IAM role for MediaLive Cluster
data "aws_iam_policy_document" "assume_role" {
  statement {
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["medialive.amazonaws.com"]
    }
    actions = ["sts:AssumeRole"]
  }
}

resource "aws_iam_role" "medialive_cluster_role" {
  name               = "MediaLiveClusterRole"
  assume_role_policy = data.aws_iam_policy_document.assume_role.json
  path               = "/"
  tags = {
    ModifiedBy = "AWSCC"
  }
}

# Create IAM policy for MediaLive Cluster
data "aws_iam_policy_document" "medialive_cluster_policy" {
  statement {
    effect = "Allow"
    actions = [
      "medialive:*",
      "ec2:CreateNetworkInterface",
      "ec2:DeleteNetworkInterface",
      "ec2:DescribeNetworkInterfaces",
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents",
      "logs:DescribeLogGroups",
      "logs:DescribeLogStreams"
    ]
    resources = ["*"]
  }
}

resource "aws_iam_role_policy" "medialive_cluster_policy" {
  name   = "MediaLiveClusterPolicy"
  role   = aws_iam_role.medialive_cluster_role.name
  policy = data.aws_iam_policy_document.medialive_cluster_policy.json
}

# Create MediaLive Cluster
resource "awscc_medialive_cluster" "example" {
  name              = "example-medialive-cluster"
  cluster_type      = "EC2"
  instance_role_arn = aws_iam_role.medialive_cluster_role.arn

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}