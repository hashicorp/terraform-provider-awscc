package cfschema

import "strings"

const (
	PropertyJsonPointerSeparator = "/"
)

// PropertyJsonPointer is a simplistic RFC 6901 handler for properties JSON Pointers.
type PropertyJsonPointer string

// EqualsPath returns true if all path parts match.
//
// This automatically handles stripping the /properties prefix.
func (p *PropertyJsonPointer) EqualsPath(path []string) bool {
	if p == nil || *p == "" {
		return false
	}

	trimmedPath := strings.TrimPrefix(string(*p), "/properties/"+strings.Join(path, PropertyJsonPointerSeparator))

	return !strings.Contains(trimmedPath, PropertyJsonPointerSeparator)
}

// EqualsStringPath returns true if the path string matches.
//
// This automatically handles stripping the /properties prefix.
func (p *PropertyJsonPointer) EqualsStringPath(path string) bool {
	if p == nil || *p == "" {
		return false
	}

	trimmedPath := strings.TrimPrefix(string(*p), "/properties")

	return trimmedPath == path
}

// Path returns the path parts.
//
// This automatically handles stripping the /properties path part.
func (p *PropertyJsonPointer) Path() []string {
	if p == nil {
		return nil
	}

	pathParts := strings.Split(strings.TrimPrefix(string(*p), "/properties/"), PropertyJsonPointerSeparator)

	return pathParts
}

// String returns a string representation of the PropertyJsonPointer.
func (p *PropertyJsonPointer) String() string {
	if p == nil {
		return ""
	}

	return string(*p)
}
