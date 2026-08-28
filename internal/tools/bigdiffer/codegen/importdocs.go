// SPDX-License-Identifier: MPL-2.0

package codegen

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-provider-awscc/internal/tools/bigdiffer/naming"
)

// Import-example doc templates, copied verbatim from the legacy import-examples
// generator. Rendered through the same parseTemplate (with the Split func), so
// the output is byte-identical to the legacy tool. The output is Terraform/shell,
// not Go, so it is not gofmt'd (docs-fmt runs terraform fmt afterward).
const (
	importExampleShellTemplate    = `$ terraform import {{ .ResourceType }}.example {{ .Identifier }}`
	importExampleStringIDTemplate = `import {
  to = {{ .ResourceType }}.example
  id = {{ .Identifier }}
}`
	importExampleIdentityTemplate = `import {
  to = {{ .ResourceType }}.example
  identity = { {{ $parts := Split .Identifier "|" }} {{ range $part := $parts }}
{{ $part }} = "{{ $part }}" {{ end }}
  }
}
`
	importExampleListResourceTemplate = `list "{{ .ResourceType }}" "example" {
  provider = awscc
}
`
)

type importExampleDocData struct {
	ResourceType string
	Identifier   string
}

// ImportExampleFile is one rendered import-example doc file: its path relative to
// the examples/ root, and its content.
type ImportExampleFile struct {
	RelPath string
	Content []byte
}

// GenerateImportExampleDocs renders the import-example doc files for one resource:
// import.sh, import-by-string-id.tf, and import-by-identity.tf under
// resources/<resource>/, plus list-resources/<resource>/list-resource.tfquery.hcl
// when the resource has a list resource. It is a faithful copy of the legacy
// import-examples generator (same templates, same identifier formatting).
func GenerateImportExampleDocs(ex ImportExample) ([]ImportExampleFile, error) {
	data := importExampleDocData{
		ResourceType: ex.ResourceName,
		Identifier:   formatImportIdentifier(ex.Identifier),
	}

	targets := []struct {
		relPath  string
		template string
	}{
		{fmt.Sprintf("resources/%s/import.sh", ex.ResourceName), importExampleShellTemplate},
		{fmt.Sprintf("resources/%s/import-by-string-id.tf", ex.ResourceName), importExampleStringIDTemplate},
		{fmt.Sprintf("resources/%s/import-by-identity.tf", ex.ResourceName), importExampleIdentityTemplate},
	}
	if ex.GenerateListResource {
		targets = append(targets, struct {
			relPath  string
			template string
		}{fmt.Sprintf("list-resources/%s/list-resource.tfquery.hcl", ex.ResourceName), importExampleListResourceTemplate})
	}

	files := make([]ImportExampleFile, 0, len(targets))
	for _, t := range targets {
		body, err := parseTemplate("import", t.template, data)
		if err != nil {
			return nil, fmt.Errorf("rendering %s: %w", t.relPath, err)
		}
		files = append(files, ImportExampleFile{RelPath: t.relPath, Content: body})
	}
	return files, nil
}

// formatImportIdentifier snake-cases each identifier segment and joins them with
// "|", quoted — matching the legacy generator exactly. An empty identifier
// renders the placeholder used by the legacy tool.
func formatImportIdentifier(identifier []string) string {
	if len(identifier) == 0 {
		return "<resource ID>"
	}
	out := make([]string, 0, len(identifier))
	for _, i := range identifier {
		out = append(out, naming.SnakeCase(i))
	}
	return fmt.Sprintf("\"%s\"", strings.Join(out, "|"))
}
