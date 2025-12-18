# Copyright IBM Corp. 2021, 2025
# SPDX-License-Identifier: MPL-2.0

# Reference: https://github.com/hashicorp/security-scanner/blob/main/CONFIG.md#binary (private repository)

binary {
  secrets {
    all = true
  }
  go_modules   = true
  osv          = true
  oss_index    = false
  nvd          = false
}