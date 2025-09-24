resource "awscc_organizations_organizational_unit" "level_1_ou" {
  name      = "child-ou-lv1"
  parent_id = var.root_id
}

resource "awscc_organizations_organizational_unit" "level_2_ou" {
  name      = "child-ou-lv2"
  parent_id = awscc_organizations_organizational_unit.level_1_ou.id
}