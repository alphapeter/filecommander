#!/bin/bash

echo using docker to build
docker run --rm -v $(pwd):/go/src/github.com/alphapeter/filecommander  -it -w /go/src/github.com/alphapeter/filecommander ubbe ./build.sh
