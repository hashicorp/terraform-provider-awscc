package acctest

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	tfcloudformation "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/service/cloudformation"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/tfresource"
)

type checkThat struct {
	testData TestData
}

func CheckThat(testData TestData) checkThat {
	return checkThat{
		testData: testData,
	}
}

// CheckDestroy returns a TestCheckFunc that tests whether a resource exists in AWS.
func (t checkThat) ExistsInAWS() resource.TestCheckFunc {
	return func(state *terraform.State) error {
		// TODO: Get the CF client from the provider.
		return existsFunc(true)(context.TODO(), nil, t.testData.CloudFormationResourceType, t.testData.TerraformResourceType, t.testData.ResourceName)(state)
	}
}

// CheckDestroy returns a TestCheckFunc that tests whether a resource has been destroyed in AWS.
func (td TestData) CheckDestroy(ctx context.Context, conn *cloudformation.Client) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		for resourceName, resourceState := range state.RootModule().Resources {
			if resourceState.Type != td.TerraformResourceType {
				continue
			}

			if resourceName != td.ResourceName {
				continue
			}

			return existsFunc(false)(ctx, conn, td.CloudFormationResourceType, td.TerraformResourceType, td.ResourceName)(state)
		}

		return nil
	}
}

func existsFunc(shouldExist bool) func(context.Context, *cloudformation.Client, string, string, string) resource.TestCheckFunc {
	return func(ctx context.Context, conn *cloudformation.Client, cfTypeName, tfTypeName, resourceName string) resource.TestCheckFunc {
		return func(state *terraform.State) error {
			rs, ok := state.RootModule().Resources[resourceName]
			if !ok {
				return fmt.Errorf("not found: %s", resourceName)
			}

			id := rs.Primary.ID

			if id == "" {
				return fmt.Errorf("no ID is set")
			}

			_, err := tfcloudformation.FindResourceByTypeNameAndID(ctx, conn, "", cfTypeName, id)

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
