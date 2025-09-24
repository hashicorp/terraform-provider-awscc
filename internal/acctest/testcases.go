// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package acctest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func (td TestData) DataSourceTest(t *testing.T, steps []resource.TestStep) {
	td.runAcceptanceTest(t, resource.TestCase{
		PreCheck:     func() { PreCheck(t) },
		CheckDestroy: nil,
		Steps:        steps,
	})
}

func (td TestData) ResourceTest(t *testing.T, steps []resource.TestStep) {
	td.runAcceptanceTest(t, resource.TestCase{
		PreCheck:     func() { PreCheck(t) },
		CheckDestroy: td.CheckDestroy(),
		Steps:        steps,
	})
}

func (td TestData) ResourceTestNoProviderFactories(t *testing.T, steps []resource.TestStep) {
	td.runAcceptanceTestNoProviderFactories(t, resource.TestCase{
		PreCheck:     func() { PreCheck(t) },
		CheckDestroy: td.CheckDestroy(),
		Steps:        steps,
	})
}

func (td TestData) runAcceptanceTest(t *testing.T, testCase resource.TestCase) {
	testCase.ProtoV6ProviderFactories = td.ProviderFactories()

	resource.ParallelTest(t, testCase)
}

func (td TestData) runAcceptanceTestNoProviderFactories(t *testing.T, testCase resource.TestCase) {
	resource.ParallelTest(t, testCase)
}

func (td TestData) ProviderFactories() map[string]func() (tfprotov6.ProviderServer, error) {
	return map[string]func() (tfprotov6.ProviderServer, error){
		"awscc": providerserver.NewProtocol6WithError(td.provider),
	}
}

func PreCheck(t *testing.T) {}
