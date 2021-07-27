package check

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/acctest/types"
)

type that struct {
	resourceName string
}

func That(resourceName string) that {
	return that{
		resourceName: resourceName,
	}
}

func (t that) ExistsInAWS(testResource types.TestResource) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		// TODO
		return nil
	}
}
