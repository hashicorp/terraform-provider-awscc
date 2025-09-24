# Create S3 bucket for storing workflow definition
resource "aws_s3_bucket" "workflow_bucket" {
  bucket = "example-omics-workflow-bucket"
}

# Create workflow file
resource "local_file" "workflow_file" {
  content  = <<EOF
version 1.0

workflow hello_world {
  call hello
}

task hello {
  command {
    echo "Hello World!"
  }
  output {
    String message = read_string(stdout())
  }
}
EOF
  filename = "${path.module}/workflow.wdl"
}

# Create ZIP file
resource "null_resource" "create_zip" {
  provisioner "local-exec" {
    command     = "zip workflow.zip workflow.wdl"
    working_dir = path.module
  }

  depends_on = [
    local_file.workflow_file
  ]
}

# Create a sample workflow definition ZIP file
resource "aws_s3_object" "workflow_definition" {
  bucket = aws_s3_bucket.workflow_bucket.id
  key    = "workflows/example-workflow.zip"
  source = "${path.module}/workflow.zip"

  depends_on = [
    null_resource.create_zip
  ]
}

# AWS Omics Workflow resource
resource "awscc_omics_workflow" "example" {
  name        = "example-workflow"
  description = "Example HealthOmics workflow created with Terraform"

  definition_uri   = "s3://${aws_s3_bucket.workflow_bucket.bucket}/${aws_s3_object.workflow_definition.key}"
  main             = "workflow.wdl"
  engine           = "WDL"
  storage_capacity = 1024

  tags = {
    Environment = "example"
    Name        = "example-workflow"
  }

  depends_on = [
    aws_s3_object.workflow_definition
  ]
}
