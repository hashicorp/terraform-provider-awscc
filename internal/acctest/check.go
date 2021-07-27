package acctest

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/acctest/types"
)

type checkThat struct {
	resourceName string
}

func CheckThat(resourceName string) checkThat {
	return checkThat{
		resourceName: resourceName,
	}
}

func (t checkThat) ExistsInAWS(testResource types.TestResource) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		// TODO
		return nil
	}
}
