#!/bin/bash
cd landing_service && go mod download 
cd ../login_service && go mod download && cd ../
go build -C landing_service
go build -C login_service