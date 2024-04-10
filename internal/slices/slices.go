// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package slices

// AppendUnique appends unique (not already in the slice) values to a slice.
func AppendUnique[S ~[]E, E comparable](s S, vs ...E) S {
	for _, v := range vs {
		var exists bool

		for _, e := range s {
			if e == v {
				exists = true
				break
			}
		}

		if !exists {
			s = append(s, v)
		}
	}

	return s
}
