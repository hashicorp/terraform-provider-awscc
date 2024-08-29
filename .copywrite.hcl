schema_version = 1

project {
  license = "MPL-2.0"

  # (OPTIONAL) A list of globs that should not have copyright/license headers.
  # Supports doublestar glob patterns for more flexibility in defining which
  # files or folders should be ignored
  header_ignore = [
    "internal/provider/generators/allschemas/*.hcl",
    "internal/provider/import_examples_gen.json",
    "examples/resources/*/import.sh",
    "examples/resources/*/*.tf",
    "examples/data-sources/*/*.tf",
    ".github/ISSUE_TEMPLATE/*.yml",
  ]
}
