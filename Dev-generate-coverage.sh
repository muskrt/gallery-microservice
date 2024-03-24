#!/bin/bash
services=`ls  | grep '_service'`
for service in $services 
do 
    cd $service 
    go test -coverprofile=coverage.txt 
    docker run --rm -v `pwd`:/app -w /app keemyb/gocover-cobertura gocover-cobertura < coverage.txt > coverage.xml
    cd ../
done 