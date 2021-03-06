#!/usr/bin/env bash
#
# This is a helper script for validating bash scripts inside the test-infra repository.
# It uses shellcheck as a validator.
set -e
set -o pipefail

readonly DEVELOPMENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
export LC_ALL=C.UTF-8
find "${DEVELOPMENT_DIR}/" -type f -name "*.sh" -exec "shellcheck" {} +
find "${DEVELOPMENT_DIR}/../prow" -type f -name "*.sh" -exec "shellcheck" {} +

echo "No issues detected!"