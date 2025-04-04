resource "awscc_sagemaker_studio_lifecycle_config" "example" {
  studio_lifecycle_config_name     = "example-studio-lc-config"
  studio_lifecycle_config_app_type = "JupyterServer"
  studio_lifecycle_config_content = base64encode(<<-EOT
    #!/bin/bash
    set -ex

    # Install Python packages
    pip install --upgrade pandas scikit-learn matplotlib seaborn

    # Print success message
    echo "Successfully installed required packages"
  EOT
  )

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}