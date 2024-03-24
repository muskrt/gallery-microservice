#!/bin/bash
echo "Running build script "
echo $PWD
go build -C ./landing_service
go build -C ./login_service