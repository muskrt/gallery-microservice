#!/bin/bash
echo "Running build script "
echo $PWD
cd landing_service 
export GOPROXY=direct

go clean --modcache
go get -u
cd ../login_service 
go get -u 
cd ../

go build -C ./landing_service
go build -C ./login_service