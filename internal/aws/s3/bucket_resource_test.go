// Copyright IBM Corp. 2021, 2025
// SPDX-License-Identifier: MPL-2.0

package s3_test

import (
	"fmt"
	"testing"

	"github.com/YakDriver/regexache"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSS3Bucket_identity_noTerraformSupport(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::S3::Bucket", "awscc_s3_bucket", "test")
	resourceName := td.ResourceName
	rName := td.RandomName()

	td.ResourceTestWithTestCase(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipAbove(tfversion.Version1_11_0),
		},
		PreCheck: func() { acctest.PreCheck(t) },
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"awscc": {
						Source:            "hashicorp/awscc",
						VersionConstraint: "1.58.0",
					},
				},
				Config: testAccAWSS3BucketConfig(&td, rName, "testing"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "tags.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.0.key", "Name"),
					resource.TestCheckResourceAttr(resourceName, "tags.0.value", rName),
					resource.TestCheckResourceAttr(resourceName, "tags.1.key", "tag"),
					resource.TestCheckResourceAttr(resourceName, "tags.1.value", "testing"),
				),
			},
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"awscc": {
						Source:            "hashicorp/awscc",
						VersionConstraint: "1.58.0",
					},
				},
				Config:      testAccAWSS3BucketConfig(&td, rName, "updated"),
				ExpectError: regexache.MustCompile(`Missing Resource Identity After Update`),
			},
			{
				Config:                   testAccAWSS3BucketConfig(&td, rName, "original"),
				ProtoV6ProviderFactories: td.ProviderFactories(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "tags.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.0.key", "Name"),
					resource.TestCheckResourceAttr(resourceName, "tags.0.value", rName),
					resource.TestCheckResourceAttr(resourceName, "tags.1.key", "tag"),
					resource.TestCheckResourceAttr(resourceName, "tags.1.value", "original"),
				),
			},
		},
	})
}

func testAccAWSS3BucketConfig(td *acctest.TestData, rName, tag string) string {
	return fmt.Sprintf(`
resource %[1]q %[2]q {
  bucket_name = %[3]q

  tags = [
    {
      key   = "Name"
      value = %[3]q
    },
    {
      key   = "tag"
      value = %[4]q
    }
  ]
}
`, td.TerraformResourceType, td.ResourceLabel, rName, tag)
}
