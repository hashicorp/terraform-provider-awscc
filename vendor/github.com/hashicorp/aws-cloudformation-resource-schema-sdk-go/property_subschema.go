package cfschema

type PropertySubschema struct {
	AllOf    []*PropertySubschema `json:"allOf,omitempty"`
	AnyOf    []*PropertySubschema `json:"anyOf,omitempty"`
	OneOf    []*PropertySubschema `json:"oneOf,omitempty"`
	Required []string             `json:"required,omitempty"`
}
