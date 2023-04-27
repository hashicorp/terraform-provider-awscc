resource "awscc_organizations_organizational_unit" "this" {
  name      = "child-ou-lv1"
  parent_id = var.root_id
}