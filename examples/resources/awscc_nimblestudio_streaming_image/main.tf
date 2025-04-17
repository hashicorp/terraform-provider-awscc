# Create a studio first
resource "awscc_nimblestudio_studio" "example" {
  admin_role_arn = aws_iam_role.nimble_admin_role.arn
  display_name   = "ExampleStudio"
  studio_name    = "example-studio"
  studio_encryption_configuration = {
    key_type = "AWS_OWNED_KEY"
  }
  user_role_arn = aws_iam_role.nimble_user_role.arn
  tags = {
    "Modified By" = "AWSCC"
  }
}

# Create the streaming image
resource "awscc_nimblestudio_streaming_image" "example" {
  name          = "ExampleStreamingImage"
  description   = "Example streaming image for Nimble Studio"
  ec_2_image_id = "ami-0568773882d492fc8" # Example AMI ID - Windows Server 2019 Base
  studio_id     = awscc_nimblestudio_studio.example.id

  encryption_configuration_key_type = "AWS_OWNED_KEY"

  tags = {
    "Modified By" = "AWSCC"
  }
}

# Required IAM roles for NimbleStudio
resource "aws_iam_role" "nimble_admin_role" {
  name = "example-nimble-admin-role"
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "nimble.amazonaws.com"
        }
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "nimble_admin_policy" {
  role       = aws_iam_role.nimble_admin_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSNimbleStudioStudioAdmin"
}

resource "aws_iam_role" "nimble_user_role" {
  name = "example-nimble-user-role"
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "nimble.amazonaws.com"
        }
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "nimble_user_policy" {
  role       = aws_iam_role.nimble_user_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSNimbleStudioUser"
}