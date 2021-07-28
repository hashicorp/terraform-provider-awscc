package acctest

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

type DisappearsStepData struct {
	// Config is a function which returns the Terraform configuration which should be used for this step.
	Config func(data TestData) string
}

// DisappearsStep returns a TestStep which first confirms the resource exists,
// then destroys it, and expects that the plan at the end of this should show
// that the resource needs to be created (since it's been destroyed).
func (td TestData) DisappearsStep(data DisappearsStepData) resource.TestStep {
	return resource.TestStep{
		Config: data.Config(td),
		Check: resource.ComposeTestCheckFunc(
			td.CheckExistsInAWS(),
			td.DeleteResource(),
		),
		ExpectNonEmptyPlan: true,
	}
}
