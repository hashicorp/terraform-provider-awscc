// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

// Legacy code: superseded by internal/tools/bigdiffer/naming, bigdiffer's
// owned copy (which also fixes a confirmed inflection.AddIrregular data race
// this package has). Kept as the deprecated fallback, used only by the legacy
// generators and internal/update; see
// contributing/docs/generating-the-provider-with-bigdiffer.md#fallback-the-legacy-process
// and contributing/docs/removing-the-legacy-generation-process.md.

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
		"OTel":           "Otel",
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

func SnakeCase(s string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(s, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")

	return strings.ToLower(snake)
}

// Pluralize converts a name to its plural form.
// The inflection package is used as a first attempt to pluralize names,
// but exceptions to the rule are handled as follows:
//   - 's' is appended to a name ending in a number, e.g. 's3' => 's3s'
//   - '_plural' is appended to any name the inflection package leaves unchanged
//     (i.e. already plural, e.g. 'windows', 'settings', 'preferences') so the
//     plural name never collides with the singular
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

	if arr := []byte(pluralName); isNumeric(arr[len(arr)-1]) {
		return pluralName + "s" // "s3" => "s3s"
	}

	// inflection left the name unchanged and no digit rule applied: the name is
	// already plural (e.g. "preferences"), so returning it unchanged would make
	// the plural data source name identical to the singular and collide. Append
	// the plural suffix to keep the two names distinct.
	return pluralName + "_plural"
}

// PluralizeWithCustomNameSuffix converts a name to its plural form similar to Pluralize,
// with the exception that a caller-chosen suffix is appended (instead of
// '_plural') when inflection leaves the name unchanged.
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

	if arr := []byte(pluralName); isNumeric(arr[len(arr)-1]) {
		return pluralName + "s" // "s3" => "s3s"
	}

	// inflection left the name unchanged and no digit rule applied: the name is
	// already plural (e.g. "preferences"), so returning it unchanged would make
	// the plural data source name identical to the singular and collide. Append
	// the caller's suffix to keep the two names distinct.
	return pluralName + suffix
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
