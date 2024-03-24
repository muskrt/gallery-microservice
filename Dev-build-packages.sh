#!/bin/bash
echo "Running build script "
echo $PWD
export GOPROXY=direct
cd landing_service 
go clean --modcache
go get -u
cd ../login_service
go clean --modcache
go get -u
cd ../
go build -C ./landing_service
go build -C ./login_service