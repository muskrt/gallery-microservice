#!/bin/bash
services=`ls  | grep '_service'`
for service in $services 
do 
    cd $service 
     docker run --rm -v `pwd`:/app -w /app golang:1.22-alpine go test -coverprofile=coverage.txt  
    docker run --rm -v `pwd`:/app -w /app golang:1.22-alpine go tool cover -func=coverage.txt
    docker run --rm -v `pwd`:/app -w /app keemyb/gocover-cobertura gocover-cobertura < coverage.txt > coverage.xml
    cd ../
done 