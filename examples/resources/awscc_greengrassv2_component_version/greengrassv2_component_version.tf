resource "awscc_greengrassv2_component_version" "MyGreengrassComponentVersion_example" {
  inline_recipe = jsonencode({
    "RecipeFormatVersion"  = "2020-01-25"
    "ComponentName"        = "MyLambdaComponent"
    "ComponentVersion"     = "1.0.0"
    "ComponentDescription" = "This is a sample Greengrass component created using InlineRecipe."
    "Manifests" = [
      {
        "Platform" = {
          "os"   = "linux"
          "arch" = "armhf"
        }
        "Lifecycle" = {
          "install" = {
            "script" = "apt-get install -y my-package"
          }
          "run" = {
            "script" = "python3 my_script.py"
          }
        }
      }
    ]
  })

  tags = {
    MyTagKey = "MyTagValue"
  }
}