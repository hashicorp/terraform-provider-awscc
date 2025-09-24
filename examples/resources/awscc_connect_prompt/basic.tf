resource "awscc_connect_prompt" "example" {
  instance_arn = aws_connect_instance.example.arn
  name         = "example.wav"
  description  = "example prompt"
  s3_uri       = "s3://${aws_s3_object.example.bucket}/${aws_s3_object.example.key}"
}