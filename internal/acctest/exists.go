package acctest

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	hclog "github.com/hashicorp/go-hclog"
	tflog "github.com/hashicorp/terraform-plugin-log"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	tfcloudformation "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/service/cloudformation"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/tfresource"
)

// CheckDestroy returns a TestCheckFunc that tests whether a resource has been destroyed in AWS.
func (td TestData) CheckDestroy() resource.TestCheckFunc {
	return func(state *terraform.State) error {
		for resourceName, resourceState := range state.RootModule().Resources {
			if resourceState.Type != td.TerraformResourceType {
				continue
			}

			if resourceName != td.ResourceName {
				continue
			}

			return td.checkExists(false)(state)
		}

		return nil
	}
}

// CheckExistsInAWS returns a TestCheckFunc that tests whether a resource exists in AWS.
func (td TestData) CheckExistsInAWS() resource.TestCheckFunc {
	return td.checkExists(true)
}

func (td TestData) checkExists(shouldExist bool) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		provider, ok := td.provider.(tfcloudformation.Provider)
		if !ok {
			return fmt.Errorf("unable to convert %T to CloudFormationProvider", td.provider)
		}

		ctx := context.TODO()
		ctx = tflog.New(ctx, tflog.WithStderrFromInit(), tflog.WithLevel(hclog.Trace), tflog.WithoutLocation())

		return existsFunc(shouldExist)(
			ctx,
			provider.CloudFormationClient(ctx),
			provider.RoleARN(ctx),
			td.CloudFormationResourceType,
			td.TerraformResourceType,
			td.ResourceName,
		)(state)
	}
}

func existsFunc(shouldExist bool) func(context.Context, *cloudformation.Client, string, string, string, string) resource.TestCheckFunc {
	return func(ctx context.Context, conn *cloudformation.Client, roleARN, cfTypeName, tfTypeName, resourceName string) resource.TestCheckFunc {
		return func(state *terraform.State) error {
			rs, ok := state.RootModule().Resources[resourceName]
			if !ok {
				return fmt.Errorf("not found: %s", resourceName)
			}

			id := rs.Primary.ID

			if id == "" {
				return fmt.Errorf("no ID is set")
			}

			_, err := tfcloudformation.FindResourceByTypeNameAndID(ctx, conn, roleARN, cfTypeName, id)

			// TODO
			// TODO Some resource can still be found but are logically deleted.
			// TODO

			if !shouldExist {
				if err != nil {
					return fmt.Errorf("(%s/%s) resource (%s) still exists", cfTypeName, tfTypeName, id)
				}

				if tfresource.NotFound(err) {
					return nil
				}
			}

			if err != nil {
				return fmt.Errorf("error reading (%s/%s) resource (%s): %w", cfTypeName, tfTypeName, id, err)
			}

			return nil
		}
	}
}
