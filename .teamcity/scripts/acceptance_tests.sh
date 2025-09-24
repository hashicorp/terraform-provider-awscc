#!/bin/bash
# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0


echo "$GOPATH"
goenv versions
echo "go env:"
go env

# shellcheck disable=2157 # These aren't constant strings, they're TeamCity variable substitution
if [[ -n "%ACCTEST_ROLE_ARN%" ]]; then
	conf=$(pwd)/aws.conf

	function cleanup {
		rm "${conf}"
	}
	trap cleanup EXIT

	touch "${conf}"
	chmod 600 "${conf}"

	export AWS_CONFIG_FILE="${conf}"

	# shellcheck disable=2157 # This isn't a constant string, it's a TeamCity variable substitution
	if [[ -n "%ACCTEST_ROLE_ARN%" ]]; then
		cat <<EOF >>"${conf}"
[profile primary]
role_arn       = %ACCTEST_ROLE_ARN%
source_profile = primary_user

[profile primary_user]
aws_access_key_id     = %AWS_ACCESS_KEY_ID%
aws_secret_access_key = %AWS_SECRET_ACCESS_KEY%
EOF

		unset AWS_ACCESS_KEY_ID
		unset AWS_SECRET_ACCESS_KEY

		export AWS_PROFILE=primary
	fi
fi

readonly TEST_BINARY='./test-binary'

# We need to loop here, since `go test -c` only allows a single package at a time
for DIR in $(find ./internal/aws/ -type d | sort -u); do
    pushd "$DIR" || continue

    go test -c . -o $TEST_BINARY
    TEST_LIST=$($TEST_BINARY -test.list="%TEST_PATTERN%")
    TEST_COUNT=$(echo "${TEST_LIST}" | wc -l)
    echo "Running ${TEST_COUNT} tests in ${DIR}..."
    echo "${TEST_LIST}" | TF_ACC=1 teamcity-go-test -test $TEST_BINARY -parallelism %ACCTEST_PARALLELISM%

    popd || continue
done