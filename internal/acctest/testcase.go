package acctest

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/provider"
)

func (td TestData) ResourceTest(t *testing.T, steps []resource.TestStep) {
	testCase := resource.TestCase{
		PreCheck: func() { PreCheck(t) },
		// TODO: Get the CF client from the provider.
		CheckDestroy: td.CheckDestroy(context.TODO(), nil),
		Steps:        steps,
	}
	td.runAcceptanceTest(t, testCase)
}

func (td TestData) runAcceptanceTest(t *testing.T, testCase resource.TestCase) {
	testCase.ProtoV6ProviderFactories = td.providers()

	resource.ParallelTest(t, testCase)
}

func (td TestData) providers() map[string]func() (tfprotov6.ProviderServer, error) {
	return map[string]func() (tfprotov6.ProviderServer, error){
		"cloudapi": func() (tfprotov6.ProviderServer, error) {
			return tfsdk.NewProtocol6Server(provider.New()), nil
		},
	}
}
