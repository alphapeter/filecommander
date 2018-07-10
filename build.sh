#!/bin/bash

echo GOPATH: $GOPATH
echo GOVERSION: 
go version
BASE_PATH=/go/src/github.com/alphapeter/filecommander
SETTINGS_PATH=$BASE_PATH/exampleconfig
ARTIFACT_PATH=$BASE_PATH/release/artifacts
OUTPUT_PATH=$BASE_PATH/release/output

#compile frontend
cd frontend
npm install
npm run build
cd ..

# compile backend
cd server
GOOS=windows GOARCH=386 go build -o $OUTPUT_PATH/windows/x86/filecommander.exe main.go
GOOS=windows GOARCH=amd64 go build -o $OUTPUT_PATH/windows/amd64/filecommander.exe main.go
GOOS=linux GOARCH=386 go build -o $OUTPUT_PATH/linux/x86/filecommander main.go
GOOS=linux GOARCH=amd64 go build -o $OUTPUT_PATH/linux/amd64/filecommander main.go
GOOS=darwin GOARCH=amd64 go build -o $OUTPUT_PATH/osx/filecommander main.go

chmod 755 $OUTPUT_PATH/linux/x86/filecommander
chmod 755 $OUTPUT_PATH/linux/amd64/filecommander

chmod 755 $OUTPUT_PATH/osx/filecommander

cp $SETTINGS_PATH/windows/settings.json $OUTPUT_PATH/windows/x86/
cp $SETTINGS_PATH/windows/settings.json $OUTPUT_PATH/windows/amd64/

cp $SETTINGS_PATH/linux/settings.json $OUTPUT_PATH/linux/x86/
cp $SETTINGS_PATH/linux/settings.json $OUTPUT_PATH/linux/amd64/

cp $SETTINGS_PATH/linux/settings.json $OUTPUT_PATH/osx/

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

