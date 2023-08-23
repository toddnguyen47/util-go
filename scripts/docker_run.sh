#!/bin/bash

set -euxo pipefail

targetDir="/go/src/tmp-go/"

docker run --rm -it --name tmp-go \
    --mount type=bind,src=$PWD,target=$targetDir \
    --workdir $targetDir \
    golang
