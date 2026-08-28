// SPDX-License-Identifier: MPL-2.0

// Package naming is bigdiffer's own copy of the naming/pluralization logic it
// needs, owned rather than imported from internal/naming. Beyond the "copy,
// don't depend on legacy" policy, the legacy Pluralize has a data race — it calls
// inflection.AddIrregular on every invocation, mutating the library's global
// rules (and triggering its lazy global compile), which the race detector flags
// under concurrency. This copy fixes that by registering the irregular once and
// serializing all access to the (non-thread-safe) inflection package. The
// parsing/attribute logic is copied verbatim so derived names stay identical.
package naming

import (
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/jinzhu/inflection"
)

const OrganizationNameAWS = "AWS"

var cfnTypeNameRE = regexp.MustCompile(`^([a-zA-Z0-9]{2,64})::([a-zA-Z0-9]{2,64})::([a-zA-Z0-9]{2,64})$`)

// ParseCloudFormationTypeName splits e.g. "AWS::Logs::LogGroup" into org, svc, res.
func ParseCloudFormationTypeName(typeName string) (string, string, string, error) {
	m := cfnTypeNameRE.FindStringSubmatch(typeName)
	if len(m) != 4 {
		return "", "", "", fmt.Errorf("matching CloudFormation type name returned %d matches; expected 4", len(m))
	}
	return m[1], m[2], m[3], nil
}

const tfTypeNameSep = "_"

var tfTypeNameRE = regexp.MustCompile(`^([a-zA-Z0-9]{2,64})` + tfTypeNameSep + `([a-zA-Z0-9]{2,64})` + tfTypeNameSep + `([a-zA-Z0-9_]{2,})$`)

// ParseTerraformTypeName splits e.g. "aws_logs_log_group" into org, svc, res.
func ParseTerraformTypeName(typeName string) (string, string, string, error) {
	m := tfTypeNameRE.FindStringSubmatch(typeName)
	if len(m) != 4 {
		return "", "", "", fmt.Errorf("matching Terraform type name returned %d matches; expected 4", len(m))
	}
	return m[1], m[2], m[3], nil
}

// CreateTerraformTypeName joins org/svc/res into a Terraform type name.
func CreateTerraformTypeName(org, svc, res string) string {
	return strings.Join([]string{org, svc, res}, tfTypeNameSep)
}

// propertyNameReplacements are applied before snake-casing so acronym runs split
// as intended (e.g. "FSx" -> "Fsx" -> "fsx").
var propertyNameReplacements = map[string]string{
	"CloudFormation": "Cloudformation",
	"CloudFront":     "Cloudfront",
	"CloudWatch":     "Cloudwatch",
	"CNAMEs":         "Cnames",
	"FSx":            "Fsx",
	"OTel":           "Otel",
}

// CloudFormationPropertyToTerraformAttribute converts a CloudFormation property
// name to a Terraform attribute name, e.g. "GlobalReplicationGroupDescription" ->
// "global_replication_group_description".
func CloudFormationPropertyToTerraformAttribute(propertyName string) string {
	propertyName = strings.TrimSpace(propertyName)
	if propertyName == "" {
		return propertyName
	}

	for old, replacement := range propertyNameReplacements {
		propertyName = strings.ReplaceAll(propertyName, old, replacement)
	}

	var b strings.Builder
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

			if (isCap && nextIsLow) || (isLow && (nextIsCap || nextIsDig) || (isDig && (nextIsCap || nextIsLow))) {
				if isCap && nextIsLow {
					if prevIsCap := i > 0 && isCapitalLetter(propertyName[i-1]); prevIsCap {
						b.WriteByte('_')
					}
				}
				b.WriteByte(ch)
				if isLow || isDig {
					b.WriteByte('_')
				}
				continue
			}
		}

		if isCap || isLow || isDig {
			b.WriteByte(ch)
		} else {
			b.WriteByte('_')
		}
	}
	return b.String()
}

var (
	snakeFirstCapRE = regexp.MustCompile("(.)([A-Z][a-z]+)")
	snakeAllCapRE   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

// SnakeCase converts a CamelCase identifier to snake_case.
func SnakeCase(s string) string {
	snake := snakeFirstCapRE.ReplaceAllString(s, "${1}_${2}")
	snake = snakeAllCapRE.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

var (
	inflectOnce sync.Once
	inflectMu   sync.Mutex
)

// inflectPlural serializes access to jinzhu/inflection, which mutates global rule
// state (AddIrregular) and lazily compiles rules on first use — both data races
// under concurrency. Registering the irregular once and guarding every call makes
// pluralization safe to invoke from concurrent goroutines.
func inflectPlural(name string) string {
	inflectMu.Lock()
	defer inflectMu.Unlock()
	inflectOnce.Do(func() { inflection.AddIrregular("lens", "lenses") })
	return inflection.Plural(name)
}

// Pluralize converts a Terraform type name to its plural form, matching the
// legacy behavior: '_plural' for custom names, trailing 's' after a digit.
func Pluralize(name string) string {
	return pluralizeWithSuffix(name, "_plural")
}

// PluralizeWithCustomNameSuffix is Pluralize with a caller-chosen suffix for
// names that are "custom" (isCustomName), used by the plural-data-source path.
func PluralizeWithCustomNameSuffix(name, suffix string) string {
	return pluralizeWithSuffix(name, suffix)
}

func pluralizeWithSuffix(name, suffix string) string {
	if name == "" {
		return name
	}

	plural := inflectPlural(name)
	if plural != name {
		return plural
	}
	if isCustomName(plural) {
		return plural + suffix
	}
	if b := []byte(plural); isNumeric(b[len(b)-1]) {
		return plural + "s"
	}
	return plural
}

var (
	customNameRE1 = regexp.MustCompile(`((e|hd|n|z)fs|(E|HD|N|Z)FS)$`)
	customNameRE2 = regexp.MustCompile(`tions$`)
	customNameRE3 = regexp.MustCompile(`issions$`)
	customNameRE4 = regexp.MustCompile(`(W|w)indows$`)
	customNameRE5 = regexp.MustCompile(`(S|s)ettings$`)
	customNameRE6 = regexp.MustCompile(`(D|d)ata$`)
)

func isCustomName(name string) bool {
	return customNameRE1.MatchString(name) ||
		customNameRE2.MatchString(name) ||
		customNameRE3.MatchString(name) ||
		customNameRE4.MatchString(name) ||
		customNameRE5.MatchString(name) ||
		customNameRE6.MatchString(name)
}

func isCapitalLetter(ch byte) bool   { return ch >= 'A' && ch <= 'Z' }
func isLowercaseLetter(ch byte) bool { return ch >= 'a' && ch <= 'z' }
func isNumeric(ch byte) bool         { return ch >= '0' && ch <= '9' }
func toLowercaseLetter(ch byte) byte { return ch - 'A' + 'a' }
