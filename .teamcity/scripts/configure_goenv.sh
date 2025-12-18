#!/usr/bin/env bash
# Copyright IBM Corp. 2021, 2025
# SPDX-License-Identifier: MPL-2.0


set -euo pipefail

go_version="$(GOENV_GOMOD_VERSION_ENABLE=1 goenv local)"

goenv install --skip-existing "${go_version}" && goenv rehash
