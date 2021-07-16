package naming

// CloudFormationPropertyToTerraformAttribute converts a CloudFormation property name to a Terraform attribute name.
// For example `GlobalReplicationGroupDescription` -> `global_replication_group_description`.
func CloudFormationPropertyToTerraformAttribute(propertyName string) string {
	return ""
}

// TerraformAttributeToCloudFormationProperty converts a Terraform attribute name to a CloudFormation property name.
// For example `global_replication_group_description` -> `GlobalReplicationGroupDescription`.
func TerraformAttributeToCloudFormationProperty(propertyName string) string {
	return ""
}
