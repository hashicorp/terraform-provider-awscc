package types

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ attr.Type = MultisetType{}
)

type MultisetType struct {
	types.ListType
}
