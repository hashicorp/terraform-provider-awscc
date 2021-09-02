package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-provider-awscc/internal/provider"
)

// Run the docs generation tool, check its repository for more information on how it works and how docs can be customized.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

func main() {
	tfsdk.Serve(context.Background(), provider.New, tfsdk.ServeOpts{
		Name: "hashicorp/awscc",
	})
}
