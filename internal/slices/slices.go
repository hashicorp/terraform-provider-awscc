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

// ApplyToAll returns a new slice containing the results of applying the function `f` to each element of the original slice `s`.
func ApplyToAll[S ~[]E1, E1, E2 any](s S, f func(E1) E2) []E2 {
	v := make([]E2, len(s))

	for i, e := range s {
		v[i] = f(e)
	}

	return v
}
