#!/bin/bash

echo using docker to build
docker run --rm -v $(pwd):/go/src/github.com/alphapeter/filecommander -it -w /go/src/github.com/alphapeter/filecommander alphapeter/buildimage:20180710 ./build.sh

TRAVIS_TAG=latest
./zip_artifacts.sh