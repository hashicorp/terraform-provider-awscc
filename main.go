package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-provider-awscc/internal/provider"
)

func main() {
	tfsdk.Serve(context.Background(), provider.New, tfsdk.ServeOpts{
		Name: "hashicorp/awscc",
	})
}
