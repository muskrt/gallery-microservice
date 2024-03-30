#!/bin/bash
echo "Running build script "
echo $PWD
cd landing_service 
docker run --rm -v `pwd`:/app -w /app golang:1.22-alpine  /bin/sh -c "go build . "
cd ../login_service
docker run --rm -v `pwd`:/app -w /app golang:1.22-alpine /bin/sh -c "go build . "
