package naming

import (
	"regexp"
	"strings"

	"github.com/jinzhu/inflection"
)

var (
	// Replace all occurrences of these strings in the property name.
	propertyNameReplacements = map[string]string{
		"CloudFormation": "Cloudformation",
		"CloudFront":     "Cloudfront",
		"CloudWatch":     "Cloudwatch",
		"CNAMEs":         "Cnames",
		"FSx":            "Fsx",
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

// Pluralize converts a name to its plural form.
// The inflection package is used as a first attempt to pluralize names,
// but exceptions to the rule are handled as follows:
//  - '_plural' is appended to a name ending in 's' e.g. 'windows'
//  - 's' is appended to a name ending in a number
func Pluralize(name string) string {
	if name == "" {
		return name
	}

	// Custom Rule
	inflection.AddIrregular("lens", "lenses") // "lens" => "lenses"

	pluralName := inflection.Plural(name)

	if pluralName != name {
		return pluralName
	}

	if isCustomName(pluralName) {
		return pluralName + "_plural"
	}

	arr := []byte(pluralName)
	lastChar := arr[len(arr)-1]

	if isNumeric(lastChar) {
		pluralName += "s" // "s3" => "s3s"
	}

	return pluralName
}

// PluralizeWithCustomNameSuffix converts a name to its plural form similar to Pluralize,
// with the exception that a suffix can be passed in as an argument to be used
// only for names that are considered "custom" i.e. return true for isCustomName.
func PluralizeWithCustomNameSuffix(name, suffix string) string {
	if name == "" {
		return name
	}

	// Custom Rule
	inflection.AddIrregular("lens", "lenses") // "lens" => "lenses"

	pluralName := inflection.Plural(name)

	if pluralName != name {
		return pluralName
	}

	if isCustomName(pluralName) {
		return pluralName + suffix
	}

	arr := []byte(pluralName)
	lastChar := arr[len(arr)-1]

	if isNumeric(lastChar) {
		pluralName += "s" // "s3" => "s3s"
	}

	return pluralName
}

func isCapitalLetter(ch byte) bool {
	return ch >= 'A' && ch <= 'Z'
}

func isCustomName(name string) bool {
	re1 := regexp.MustCompile(`((e|hd|n|z)fs|(E|HD|N|Z)FS)$`)
	re2 := regexp.MustCompile(`tions$`)
	re3 := regexp.MustCompile(`(W|w)indows$`)

	return re1.MatchString(name) || re2.MatchString(name) || re3.MatchString(name)
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
