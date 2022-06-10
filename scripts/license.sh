#!/usr/bin/env bash
# Copyright 2021-present Airheart, Inc. All rights reserved.
# This source code is licensed under the Apache 2.0 license found
# in the LICENSE file in the root directory of this source tree.


set -euo pipefail
[[ ${DEBUG:-} ]] && set -x

# go-licenser -licensor "Airheart, Inc." -license ASL2-Short
addlicense -c "Airheart, Inc." -f .github/LICENSE_SHORT "./**/*.go"
