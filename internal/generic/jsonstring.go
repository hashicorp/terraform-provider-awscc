// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package generic

import (
	"encoding/json"
)

func expandJSONFromString(s string) (interface{}, error) {
	var v interface{}

	err := json.Unmarshal([]byte(s), &v)

	return v, err
}
