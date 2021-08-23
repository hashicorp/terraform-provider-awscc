package naming

import (
	"strings"
)

var (
	// Replace all occuurences of these strings in the property name.
	propertyNameReplacements = map[string]string{
		"CloudFormation": "Cloudformation",
		"CloudFront":     "Cloudfront",
		"CloudWatch":     "Cloudwatch",
	}
)

// CloudFormationPropertyToTerraformAttribute converts a CloudFormation property name to a Terraform attribute name.
// For example `GlobalReplicationGroupDescription` -> `global_replication_group_description`.
func CloudFormationPropertyToTerraformAttribute(propertyName string) string {
	propertyName = strings.TrimSpace(propertyName)

	if propertyName == "" {
		return propertyName
	}

	for old, new := range propertyNameReplacements {
		propertyName = strings.ReplaceAll(propertyName, old, new)
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
