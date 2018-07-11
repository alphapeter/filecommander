#!/bin/bash

echo using docker to build
docker run --rm -v $(pwd):/go/src/github.com/alphapeter/filecommander -it -w /go/src/github.com/alphapeter/filecommander alphapeter/buildimage:20180710 ./build.sh

OUTPUT_PATH=$TRAVIS_BUILD_DIR/release/output
ARTIFACT_PATH=$TRAVIS_BUILD_DIR/release/artifacts/

mkdir $ARTIFACT_PATH

cd $OUTPUT_PATH/windows/x86/
zip $ARTIFACT_PATH/filecommander-$TRAVIS_TAG-windows-32bit.zip ./*

cd $OUTPUT_PATH/windows/amd64
zip $ARTIFACT_PATH/filecommander-$TRAVIS_TAG-windows-64bit.zip ./*

cd $OUTPUT_PATH/linux/x86
tar -cvzf $ARTIFACT_PATH/filecommander-$TRAVIS_TAG-linux-x86.tar.gz ./*
cd $OUTPUT_PATH/linux/amd64
tar -cvzf $ARTIFACT_PATH/filecommander-$TRAVIS_TAG-linux-amd64.tar.gz ./*

cd $OUTPUT_PATH/osx
zip $ARTIFACT_PATH/filecommander-$TRAVIS_TAG-osx.zip ./*
