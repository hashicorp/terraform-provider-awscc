resource "awscc_connect_hours_of_operation" "example" {
  instance_arn = aws_connect_instance.example.arn
  name         = "Office hours example"
  description  = "Monday and Tuesday hours"
  time_zone    = "EST"

  config = [
    {
      day = "MONDAY"
      start_time = {
        hours   = 08
        minutes = 00
      }
      end_time = {
        hours   = 18
        minutes = 30
      }
    },
    {
      day = "TUESDAY"
      start_time = {
        hours   = 07
        minutes = 12
      }
      end_time = {
        hours   = 19
        minutes = 35
      }
    }
  ]
}