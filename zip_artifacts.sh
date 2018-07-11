#!/bin/bash

OUTPUT_PATH=$(pwd)/release/output
ARTIFACT_PATH=$(pwd)/release/artifacts
echo Creating artifacts for version $RELEASE_VERSION

mkdir -p $ARTIFACT_PATH

cd $OUTPUT_PATH/windows/x86/
zip $ARTIFACT_PATH/filecommander-$RELEASE_VERSION-windows-32bit.zip ./*

cd $OUTPUT_PATH/windows/amd64
zip $ARTIFACT_PATH/filecommander-$RELEASE_VERSION-windows-64bit.zip ./*

cd $OUTPUT_PATH/linux/x86
tar -cvzf $ARTIFACT_PATH/filecommander-$RELEASE_VERSION-linux-x86.tar.gz ./*
cd $OUTPUT_PATH/linux/amd64
tar -cvzf $ARTIFACT_PATH/filecommander-$RELEASE_VERSION-linux-amd64.tar.gz ./*

cd $OUTPUT_PATH/osx
zip $ARTIFACT_PATH/filecommander-$RELEASE_VERSION-osx.zip ./*
