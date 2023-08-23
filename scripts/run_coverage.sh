#!/bin/bash

set -euxo pipefail
#tmpDir="tmp"

#mkdir -p "${tmpDir}"
go test -v -tags dynamic -coverprofile cover.out "$1"
go tool cover -html=cover.out
