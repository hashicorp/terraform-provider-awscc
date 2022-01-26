package acctest

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-log/tfsdklog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	tfcloudcontrol "github.com/hashicorp/terraform-provider-awscc/internal/service/cloudcontrol"
	"github.com/hashicorp/terraform-provider-awscc/internal/tfresource"
)

const (
	deleteResourceTimeout = 120 * time.Minute
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

// DeleteResource returns a TestCheckFunc that deletes a resource in AWS.
func (td TestData) DeleteResource() resource.TestCheckFunc {
	return func(state *terraform.State) error {
		resourceName := td.ResourceName
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}

		id := rs.Primary.ID

		if id == "" {
			return fmt.Errorf("no ID is set")
		}

		provider, ok := td.provider.(tfcloudcontrol.Provider)
		if !ok {
			return fmt.Errorf("unable to convert %T to CloudControlApiProvider", td.provider)
		}

		ctx := getTestContext()

		return tfcloudcontrol.DeleteResource(ctx, provider.CloudControlApiClient(ctx), provider.RoleARN(ctx), td.CloudFormationResourceType, id, deleteResourceTimeout)
	}
}

func (td TestData) checkExists(shouldExist bool) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		provider, ok := td.provider.(tfcloudcontrol.Provider)
		if !ok {
			return fmt.Errorf("unable to convert %T to CloudControlApiProvider", td.provider)
		}

		ctx := getTestContext()

		return existsFunc(shouldExist)(
			ctx,
			provider.CloudControlApiClient(ctx),
			provider.RoleARN(ctx),
			td.CloudFormationResourceType,
			td.TerraformResourceType,
			td.ResourceName,
		)(state)
	}
}

func existsFunc(shouldExist bool) func(context.Context, *cloudcontrol.Client, string, string, string, string) resource.TestCheckFunc {
	return func(ctx context.Context, conn *cloudcontrol.Client, roleARN, cfTypeName, tfTypeName, resourceName string) resource.TestCheckFunc {
		return func(state *terraform.State) error {
			rs, ok := state.RootModule().Resources[resourceName]
			if !ok {
				return fmt.Errorf("not found: %s", resourceName)
			}

			id := rs.Primary.ID

			if id == "" {
				return fmt.Errorf("no ID is set")
			}

			_, err := tfcloudcontrol.FindResourceByTypeNameAndID(ctx, conn, roleARN, cfTypeName, id)

			// TODO
			// TODO Some resource can still be found but are logically deleted.
			// TODO

			if !shouldExist {
				if err == nil {
					return fmt.Errorf("(%s/%s) resource (%s) still exists", cfTypeName, tfTypeName, id)
				}

				if tfresource.NotFound(err) {
					return nil
				}
			}

			if err != nil {
				return fmt.Errorf("reading (%s/%s) resource (%s): %w", cfTypeName, tfTypeName, id, err)
			}

			return nil
		}
	}
}

func getTestContext() context.Context {
	return tfsdklog.NewRootProviderLogger(context.TODO(), tflog.WithLevelFromEnv("TF_LOG"), tflog.WithoutLocation())
}
