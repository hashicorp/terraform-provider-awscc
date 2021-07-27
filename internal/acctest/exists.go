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
		p, err := t.testData.cfProvider()

		if err != nil {
			return err
		}

		ctx := context.TODO()

		return existsFunc(true)(
			ctx,
			p.CloudFormationClient(ctx),
			p.RoleARN(ctx),
			t.testData.CloudFormationResourceType,
			t.testData.TerraformResourceType,
			t.testData.ResourceName,
		)(state)
	}
}

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

			p, err := td.cfProvider()

			if err != nil {
				return err
			}

			ctx := context.TODO()

			return existsFunc(false)(
				ctx,
				p.CloudFormationClient(ctx),
				p.RoleARN(ctx),
				td.CloudFormationResourceType,
				td.TerraformResourceType,
				td.ResourceName,
			)(state)
		}

		return nil
	}
}

func (td TestData) cfProvider() (tfcloudformation.Provider, error) {
	if provider, ok := td.provider.(tfcloudformation.Provider); ok {
		return provider, nil
	}

	return nil, fmt.Errorf("unable to convert %T to CloudFormationProvider", td.provider)
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
