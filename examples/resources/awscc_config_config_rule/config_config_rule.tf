resource "awscc_config_config_rule" "r" {
  config_rule_name = ""
  description = "Rule to validate if versioning is enabled"
  evaluation_modes = [{
    mode= "DETECTIVE"
  }]

  source = {
    owner = "AWS"
    source_identifier = "S3_BUCKET_VERSIONING_ENABLED"
    custom_policy_details ={
      policy_runtime = "guard-2.x.x"
      enable_debug_log_delivery = false
      policy_text = <<EOF
      {
          "Version": "2012-10-17",
          "id": "abc"
          "Statement": [
              {
                  "Sid": "policyforconfigrule",
                  "Effect": "Allow",
                  "Principal": {
                      "AWS": [
                          "config.amazonaws.com"
                      ]
                  },
                  "Action": [
                      "config:Put*"
                  ],
                  "Resource": "*"
              }
          ]
      }
      EOF                    
  }
   } 
}