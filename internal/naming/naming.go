package naming

// CloudFormationPropertyNameToTerraformAttributeName converts a CloudFormation property name to a Terraform attribute name.
// For example `GlobalReplicationGroupDescription` -> `global_replication_group_description`.
func CloudFormationPropertyNameToTerraformAttributeName(propertyName string) string {
	return ""
}

// TerraformAttributeNameToCloudFormationPropertyName converts a Terraform attribute name to a CloudFormation property name.
// For example `global_replication_group_description` -> `GlobalReplicationGroupDescription`.
func TerraformAttributeNameToCloudFormationPropertyName(propertyName string) string {
	return ""
}
