#!/usr/bin/env bash

echo "E2e test of netdoctor"

# Install ginkgo
GO111MODULE=on go install github.com/onsi/ginkgo/v2/ginkgo

set +e
ginkgo -v --race --trace --fail-fast -p --randomize-all ./test/e2e/ --
TESTING_RESULT=$?

exit $TESTING_RESULT