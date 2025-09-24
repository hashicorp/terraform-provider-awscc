// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package strings

import (
	"strings"
)

const (
	and      = " and "
	comma    = ", "
	commaAnd = ", and "

	lenAnd      = len(and)
	lenComma    = len(comma)
	lenCommaAnd = len(commaAnd)
)

func ProseJoin(elems []string) string {
	switch len(elems) {
	case 0:
		return ""

	case 1:
		return elems[0]

	case 2: //nolint:mnd
		size := len(elems[0]) + lenAnd + len(elems[1])
		var b strings.Builder
		b.Grow(size)
		b.WriteString(elems[0])
		b.WriteString(and)
		b.WriteString(elems[1])
		return b.String()
	}

	size := (len(elems)-2)*lenComma + lenCommaAnd //nolint:mnd
	for i := 0; i < len(elems); i++ {
		size += len(elems[i])
	}
	var b strings.Builder
	b.Grow(size)
	b.WriteString(elems[0])
	for _, s := range elems[1 : len(elems)-1] {
		b.WriteString(comma)
		b.WriteString(s)
	}
	b.WriteString(commaAnd)
	b.WriteString(elems[len(elems)-1])
	return b.String()
}
