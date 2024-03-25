#!/bin/bash
echo "Running build script "
echo $PWD
cd landing_service 
docker run --rm -v `pwd`:/app -w /app golang:1.22  /bin/sh -c "go mod download && go get -u && go build . "
cd ../login_service
docker run --rm -v `pwd`:/app -w /app golang:1.22 /bin/sh -c "go mod download && go get -u && go build . "
