# Get current account ID
data "aws_caller_identity" "current" {}

# Create QuickSight theme
resource "awscc_quicksight_theme" "example" {
  aws_account_id = data.aws_caller_identity.current.account_id
  theme_id       = "example-theme"
  name           = "Example Theme"
  base_theme_id  = "CLASSIC" # Using CLASSIC as base theme

  # Theme configuration
  configuration = {
    data_color_palette = {
      colors = [
        "#1F77B4",
        "#FF7F0E",
        "#2CA02C",
        "#D62728",
        "#9467BD",
        "#8C564B",
        "#E377C2",
        "#7F7F7F"
      ]
      empty_fill_color = "#FFFFFF"
      min_max_gradient = ["#DE4968", "#4EACC4"]
    }

    ui_color_palette = {
      primary_background   = "#FFFFFF"
      primary_foreground   = "#000000"
      secondary_background = "#F2F2F2"
      secondary_foreground = "#333333"
      accent               = "#0073BB"
      accent_foreground    = "#FFFFFF"
      danger               = "#E65100"
      danger_foreground    = "#FFFFFF"
      warning              = "#F2C94C"
      warning_foreground   = "#000000"
      success              = "#2D8700"
      success_foreground   = "#FFFFFF"
      dimension            = "#0073BB"
      dimension_foreground = "#FFFFFF"
      measure              = "#FF6B00"
      measure_foreground   = "#FFFFFF"
    }

    sheet = {
      tile = {
        border = {
          show = true
        }
      }
      tile_layout = {
        gutter = {
          show = true
        }
        margin = {
          show = true
        }
      }
    }

    typography = {
      font_families = [
        {
          font_family = "sans-serif"
        }
      ]
    }
  }

  # Optional version description
  version_description = "Initial version"

  # Tags
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}