package naming

import (
	"strings"
)

// CloudFormationPropertyToTerraformAttribute converts a CloudFormation property name to a Terraform attribute name.
// For example `GlobalReplicationGroupDescription` -> `global_replication_group_description`.
func CloudFormationPropertyToTerraformAttribute(propertyName string) string {
	propertyName = strings.TrimSpace(propertyName)

	if propertyName == "" {
		return propertyName
	}

	attributeName := strings.Builder{}

	for i, ch := range []byte(propertyName) {
		isCap := isCapitalLetter(ch)
		isLow := isLowercaseLetter(ch)
		isDig := isNumeric(ch)

		if isCap {
			ch = toLowercaseLetter(ch)
		}

		if i < len(propertyName)-1 {
			nextCh := propertyName[i+1]
			nextIsCap := isCapitalLetter(nextCh)
			nextIsLow := isLowercaseLetter(nextCh)
			nextIsDig := isNumeric(nextCh)

			// Append underscore if case changes.
			if (isCap && nextIsLow) || (isLow && (nextIsCap || nextIsDig) || (isDig && (nextIsCap || nextIsLow))) {
				if isCap && nextIsLow {
					if prevIsCap := i > 0 && isCapitalLetter(propertyName[i-1]); prevIsCap {
						attributeName.WriteByte('_')
					}
				}
				attributeName.WriteByte(ch)
				if isLow || isDig {
					attributeName.WriteByte('_')
				}

				continue
			}
		}

		if isCap || isLow || isDig {
			attributeName.WriteByte(ch)
		} else {
			attributeName.WriteByte('_')
		}
	}

	return attributeName.String()
}

// TerraformAttributeToCloudFormationProperty converts a Terraform attribute name to a CloudFormation property name.
// For example `global_replication_group_description` -> `GlobalReplicationGroupDescription`.
func TerraformAttributeToCloudFormationProperty(attributeName string) string {
	attributeName = strings.TrimSpace(attributeName)

	if attributeName == "" {
		return attributeName
	}

	capNext := true // Start with a capital letter.
	propertyName := strings.Builder{}

	for _, ch := range []byte(attributeName) {
		isCap := isCapitalLetter(ch)
		isLow := isLowercaseLetter(ch)
		isDig := isNumeric(ch)

		if capNext {
			if isLow {
				ch = toUppercaseLetter(ch)
			}
		}
		if isCap || isLow {
			propertyName.WriteByte(ch)
			capNext = false
		} else if isDig {
			propertyName.WriteByte(ch)
			capNext = true
		} else {
			capNext = true
		}
	}

	return propertyName.String()
}

func isCapitalLetter(ch byte) bool {
	return ch >= 'A' && ch <= 'Z'
}

func isLowercaseLetter(ch byte) bool {
	return ch >= 'a' && ch <= 'z'
}

func isNumeric(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func toLowercaseLetter(ch byte) byte {
	ch += 'a'
	ch -= 'A'
	return ch
}

func toUppercaseLetter(ch byte) byte {
	ch += 'A'
	ch -= 'a'
	return ch
}
