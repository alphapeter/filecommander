#!/bin/bash

echo using docker to build
docker run --rm -v $(pwd):/go/src/github.com/alphapeter/filecommander -it -w /go/src/github.com/alphapeter/filecommander -e "RELEASE_VERSION=latest" alphapeter/buildimage:20180710 bash -c "./build.sh && ./zip_artifacts.sh"

