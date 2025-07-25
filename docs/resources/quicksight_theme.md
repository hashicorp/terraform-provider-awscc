
---
page_title: "awscc_quicksight_theme Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Definition of the AWS::QuickSight::Theme Resource Type.
---

# awscc_quicksight_theme (Resource)

Definition of the AWS::QuickSight::Theme Resource Type.

## Example Usage

### Custom QuickSight Theme Configuration

Creates a QuickSight theme with customized color palettes, typography, and sheet layout settings based on the CLASSIC theme, allowing for comprehensive visual customization of QuickSight dashboards.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
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
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `aws_account_id` (String)
- `base_theme_id` (String)
- `configuration` (Attributes) <p>The theme configuration. This configuration contains all of the display properties for
            a theme.</p> (see [below for nested schema](#nestedatt--configuration))
- `name` (String)
- `theme_id` (String)

### Optional

- `permissions` (Attributes List) (see [below for nested schema](#nestedatt--permissions))
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))
- `version_description` (String)

### Read-Only

- `arn` (String) <p>The Amazon Resource Name (ARN) of the theme.</p>
- `created_time` (String) <p>The date and time that the theme was created.</p>
- `id` (String) Uniquely identifies the resource.
- `last_updated_time` (String) <p>The date and time that the theme was last updated.</p>
- `type` (String)
- `version` (Attributes) <p>A version of a theme.</p> (see [below for nested schema](#nestedatt--version))

<a id="nestedatt--configuration"></a>
### Nested Schema for `configuration`

Optional:

- `data_color_palette` (Attributes) <p>The theme colors that are used for data colors in charts. The colors description is a
            hexadecimal color code that consists of six alphanumerical characters, prefixed with
                <code>#</code>, for example #37BFF5. </p> (see [below for nested schema](#nestedatt--configuration--data_color_palette))
- `sheet` (Attributes) <p>The theme display options for sheets. </p> (see [below for nested schema](#nestedatt--configuration--sheet))
- `typography` (Attributes) (see [below for nested schema](#nestedatt--configuration--typography))
- `ui_color_palette` (Attributes) <p>The theme colors that apply to UI and to charts, excluding data colors. The colors
            description is a hexadecimal color code that consists of six alphanumerical characters,
            prefixed with <code>#</code>, for example #37BFF5. For more information, see <a href="https://docs.aws.amazon.com/quicksight/latest/user/themes-in-quicksight.html">Using Themes in Amazon QuickSight</a> in the <i>Amazon QuickSight User
                Guide.</i>
         </p> (see [below for nested schema](#nestedatt--configuration--ui_color_palette))

<a id="nestedatt--configuration--data_color_palette"></a>
### Nested Schema for `configuration.data_color_palette`

Optional:

- `colors` (List of String) <p>The hexadecimal codes for the colors.</p>
- `empty_fill_color` (String) <p>The hexadecimal code of a color that applies to charts where a lack of data is
            highlighted.</p>
- `min_max_gradient` (List of String) <p>The minimum and maximum hexadecimal codes that describe a color gradient. </p>


<a id="nestedatt--configuration--sheet"></a>
### Nested Schema for `configuration.sheet`

Optional:

- `tile` (Attributes) <p>Display options related to tiles on a sheet.</p> (see [below for nested schema](#nestedatt--configuration--sheet--tile))
- `tile_layout` (Attributes) <p>The display options for the layout of tiles on a sheet.</p> (see [below for nested schema](#nestedatt--configuration--sheet--tile_layout))

<a id="nestedatt--configuration--sheet--tile"></a>
### Nested Schema for `configuration.sheet.tile`

Optional:

- `border` (Attributes) <p>The display options for tile borders for visuals.</p> (see [below for nested schema](#nestedatt--configuration--sheet--tile--border))

<a id="nestedatt--configuration--sheet--tile--border"></a>
### Nested Schema for `configuration.sheet.tile.border`

Optional:

- `show` (Boolean) <p>The option to enable display of borders for visuals.</p>



<a id="nestedatt--configuration--sheet--tile_layout"></a>
### Nested Schema for `configuration.sheet.tile_layout`

Optional:

- `gutter` (Attributes) <p>The display options for gutter spacing between tiles on a sheet.</p> (see [below for nested schema](#nestedatt--configuration--sheet--tile_layout--gutter))
- `margin` (Attributes) <p>The display options for margins around the outside edge of sheets.</p> (see [below for nested schema](#nestedatt--configuration--sheet--tile_layout--margin))

<a id="nestedatt--configuration--sheet--tile_layout--gutter"></a>
### Nested Schema for `configuration.sheet.tile_layout.gutter`

Optional:

- `show` (Boolean) <p>This Boolean value controls whether to display a gutter space between sheet tiles.
        </p>


<a id="nestedatt--configuration--sheet--tile_layout--margin"></a>
### Nested Schema for `configuration.sheet.tile_layout.margin`

Optional:

- `show` (Boolean) <p>This Boolean value controls whether to display sheet margins.</p>




<a id="nestedatt--configuration--typography"></a>
### Nested Schema for `configuration.typography`

Optional:

- `font_families` (Attributes List) (see [below for nested schema](#nestedatt--configuration--typography--font_families))

<a id="nestedatt--configuration--typography--font_families"></a>
### Nested Schema for `configuration.typography.font_families`

Optional:

- `font_family` (String)



<a id="nestedatt--configuration--ui_color_palette"></a>
### Nested Schema for `configuration.ui_color_palette`

Optional:

- `accent` (String) <p>This color is that applies to selected states and buttons.</p>
- `accent_foreground` (String) <p>The foreground color that applies to any text or other elements that appear over the
            accent color.</p>
- `danger` (String) <p>The color that applies to error messages.</p>
- `danger_foreground` (String) <p>The foreground color that applies to any text or other elements that appear over the
            error color.</p>
- `dimension` (String) <p>The color that applies to the names of fields that are identified as
            dimensions.</p>
- `dimension_foreground` (String) <p>The foreground color that applies to any text or other elements that appear over the
            dimension color.</p>
- `measure` (String) <p>The color that applies to the names of fields that are identified as measures.</p>
- `measure_foreground` (String) <p>The foreground color that applies to any text or other elements that appear over the
            measure color.</p>
- `primary_background` (String) <p>The background color that applies to visuals and other high emphasis UI.</p>
- `primary_foreground` (String) <p>The color of text and other foreground elements that appear over the primary
            background regions, such as grid lines, borders, table banding, icons, and so on.</p>
- `secondary_background` (String) <p>The background color that applies to the sheet background and sheet controls.</p>
- `secondary_foreground` (String) <p>The foreground color that applies to any sheet title, sheet control text, or UI that
            appears over the secondary background.</p>
- `success` (String) <p>The color that applies to success messages, for example the check mark for a
            successful download.</p>
- `success_foreground` (String) <p>The foreground color that applies to any text or other elements that appear over the
            success color.</p>
- `warning` (String) <p>This color that applies to warning and informational messages.</p>
- `warning_foreground` (String) <p>The foreground color that applies to any text or other elements that appear over the
            warning color.</p>



<a id="nestedatt--permissions"></a>
### Nested Schema for `permissions`

Optional:

- `actions` (List of String) <p>The IAM action to grant or revoke permissions on.</p>
- `principal` (String) <p>The Amazon Resource Name (ARN) of the principal. This can be one of the
            following:</p>
         <ul>
            <li>
               <p>The ARN of an Amazon QuickSight user or group associated with a data source or dataset. (This is common.)</p>
            </li>
            <li>
               <p>The ARN of an Amazon QuickSight user, group, or namespace associated with an analysis, dashboard, template, or theme. (This is common.)</p>
            </li>
            <li>
               <p>The ARN of an Amazon Web Services account root: This is an IAM ARN rather than a QuickSight
                    ARN. Use this option only to share resources (templates) across Amazon Web Services accounts.
                    (This is less common.) </p>
            </li>
         </ul>


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) <p>Tag key.</p>
- `value` (String) <p>Tag value.</p>


<a id="nestedatt--version"></a>
### Nested Schema for `version`

Read-Only:

- `arn` (String) <p>The Amazon Resource Name (ARN) of the resource.</p>
- `base_theme_id` (String) <p>The Amazon QuickSight-defined ID of the theme that a custom theme inherits from. All
            themes initially inherit from a default Amazon QuickSight theme.</p>
- `configuration` (Attributes) <p>The theme configuration. This configuration contains all of the display properties for
            a theme.</p> (see [below for nested schema](#nestedatt--version--configuration))
- `created_time` (String) <p>The date and time that this theme version was created.</p>
- `description` (String) <p>The description of the theme.</p>
- `errors` (Attributes List) <p>Errors associated with the theme.</p> (see [below for nested schema](#nestedatt--version--errors))
- `status` (String)
- `version_number` (Number) <p>The version number of the theme.</p>

<a id="nestedatt--version--configuration"></a>
### Nested Schema for `version.configuration`

Read-Only:

- `data_color_palette` (Attributes) <p>The theme colors that are used for data colors in charts. The colors description is a
            hexadecimal color code that consists of six alphanumerical characters, prefixed with
                <code>#</code>, for example #37BFF5. </p> (see [below for nested schema](#nestedatt--version--configuration--data_color_palette))
- `sheet` (Attributes) <p>The theme display options for sheets. </p> (see [below for nested schema](#nestedatt--version--configuration--sheet))
- `typography` (Attributes) (see [below for nested schema](#nestedatt--version--configuration--typography))
- `ui_color_palette` (Attributes) <p>The theme colors that apply to UI and to charts, excluding data colors. The colors
            description is a hexadecimal color code that consists of six alphanumerical characters,
            prefixed with <code>#</code>, for example #37BFF5. For more information, see <a href="https://docs.aws.amazon.com/quicksight/latest/user/themes-in-quicksight.html">Using Themes in Amazon QuickSight</a> in the <i>Amazon QuickSight User
                Guide.</i>
         </p> (see [below for nested schema](#nestedatt--version--configuration--ui_color_palette))

<a id="nestedatt--version--configuration--data_color_palette"></a>
### Nested Schema for `version.configuration.data_color_palette`

Read-Only:

- `colors` (List of String) <p>The hexadecimal codes for the colors.</p>
- `empty_fill_color` (String) <p>The hexadecimal code of a color that applies to charts where a lack of data is
            highlighted.</p>
- `min_max_gradient` (List of String) <p>The minimum and maximum hexadecimal codes that describe a color gradient. </p>


<a id="nestedatt--version--configuration--sheet"></a>
### Nested Schema for `version.configuration.sheet`

Read-Only:

- `tile` (Attributes) <p>Display options related to tiles on a sheet.</p> (see [below for nested schema](#nestedatt--version--configuration--sheet--tile))
- `tile_layout` (Attributes) <p>The display options for the layout of tiles on a sheet.</p> (see [below for nested schema](#nestedatt--version--configuration--sheet--tile_layout))

<a id="nestedatt--version--configuration--sheet--tile"></a>
### Nested Schema for `version.configuration.sheet.tile`

Read-Only:

- `border` (Attributes) <p>The display options for tile borders for visuals.</p> (see [below for nested schema](#nestedatt--version--configuration--sheet--tile--border))

<a id="nestedatt--version--configuration--sheet--tile--border"></a>
### Nested Schema for `version.configuration.sheet.tile.border`

Read-Only:

- `show` (Boolean) <p>The option to enable display of borders for visuals.</p>



<a id="nestedatt--version--configuration--sheet--tile_layout"></a>
### Nested Schema for `version.configuration.sheet.tile_layout`

Read-Only:

- `gutter` (Attributes) <p>The display options for gutter spacing between tiles on a sheet.</p> (see [below for nested schema](#nestedatt--version--configuration--sheet--tile_layout--gutter))
- `margin` (Attributes) <p>The display options for margins around the outside edge of sheets.</p> (see [below for nested schema](#nestedatt--version--configuration--sheet--tile_layout--margin))

<a id="nestedatt--version--configuration--sheet--tile_layout--gutter"></a>
### Nested Schema for `version.configuration.sheet.tile_layout.gutter`

Read-Only:

- `show` (Boolean) <p>This Boolean value controls whether to display a gutter space between sheet tiles.
        </p>


<a id="nestedatt--version--configuration--sheet--tile_layout--margin"></a>
### Nested Schema for `version.configuration.sheet.tile_layout.margin`

Read-Only:

- `show` (Boolean) <p>This Boolean value controls whether to display sheet margins.</p>




<a id="nestedatt--version--configuration--typography"></a>
### Nested Schema for `version.configuration.typography`

Read-Only:

- `font_families` (Attributes List) (see [below for nested schema](#nestedatt--version--configuration--typography--font_families))

<a id="nestedatt--version--configuration--typography--font_families"></a>
### Nested Schema for `version.configuration.typography.font_families`

Read-Only:

- `font_family` (String)



<a id="nestedatt--version--configuration--ui_color_palette"></a>
### Nested Schema for `version.configuration.ui_color_palette`

Read-Only:

- `accent` (String) <p>This color is that applies to selected states and buttons.</p>
- `accent_foreground` (String) <p>The foreground color that applies to any text or other elements that appear over the
            accent color.</p>
- `danger` (String) <p>The color that applies to error messages.</p>
- `danger_foreground` (String) <p>The foreground color that applies to any text or other elements that appear over the
            error color.</p>
- `dimension` (String) <p>The color that applies to the names of fields that are identified as
            dimensions.</p>
- `dimension_foreground` (String) <p>The foreground color that applies to any text or other elements that appear over the
            dimension color.</p>
- `measure` (String) <p>The color that applies to the names of fields that are identified as measures.</p>
- `measure_foreground` (String) <p>The foreground color that applies to any text or other elements that appear over the
            measure color.</p>
- `primary_background` (String) <p>The background color that applies to visuals and other high emphasis UI.</p>
- `primary_foreground` (String) <p>The color of text and other foreground elements that appear over the primary
            background regions, such as grid lines, borders, table banding, icons, and so on.</p>
- `secondary_background` (String) <p>The background color that applies to the sheet background and sheet controls.</p>
- `secondary_foreground` (String) <p>The foreground color that applies to any sheet title, sheet control text, or UI that
            appears over the secondary background.</p>
- `success` (String) <p>The color that applies to success messages, for example the check mark for a
            successful download.</p>
- `success_foreground` (String) <p>The foreground color that applies to any text or other elements that appear over the
            success color.</p>
- `warning` (String) <p>This color that applies to warning and informational messages.</p>
- `warning_foreground` (String) <p>The foreground color that applies to any text or other elements that appear over the
            warning color.</p>



<a id="nestedatt--version--errors"></a>
### Nested Schema for `version.errors`

Read-Only:

- `message` (String) <p>The error message.</p>
- `type` (String)

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_quicksight_theme.example
  id = "theme_id|aws_account_id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_quicksight_theme.example "theme_id|aws_account_id"
```
