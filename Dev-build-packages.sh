#!/bin/bash
echo "Running build script "
echo $PWD
cd landing_service 
go mod download 
go mod tidy
cd ../login_service
go mod download 
go mod tidy
cd ../
go build -C ./landing_service
go build -C ./login_service