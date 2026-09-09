// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codegen

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// parseTemplate parses templateBody under templateName and executes it
// against templateData, returning the rendered bytes. It supports two
// template functions beyond the standard library's: Title (title-cases a
// string, used for Go identifier construction) and Split (splits a quoted
// string literal on sep, unquoting it first).
func parseTemplate(templateName, templateBody string, templateData any) ([]byte, error) {
	funcMap := template.FuncMap{
		"Title": cases.Title(language.Und, cases.NoLower).String,
		"Split": func(s, sep string) []string {
			s, _ = strconv.Unquote(s)
			return strings.Split(s, sep)
		},
	}
	tmpl, err := template.New(templateName).Funcs(funcMap).Parse(templateBody)
	if err != nil {
		return nil, fmt.Errorf("parsing function template: %w", err)
	}

	var buffer bytes.Buffer
	if err := tmpl.Execute(&buffer, templateData); err != nil {
		return nil, fmt.Errorf("executing template: %w", err)
	}

	return buffer.Bytes(), nil
}
