#!/bin/bash
services=`ls  | grep '_service'`
echo "running coverage script"
for service in $services 
do 
    cd $service 
     docker run --rm -v `pwd`:/app -w /app golang:1.22-alpine /bin/sh -c "go test -coverprofile=cover.txt  "
    docker run --rm -v `pwd`:/app -w /app golang:1.22-alpine /bin/sh -c "go tool cover -func=cover.txt"
    docker run --rm -v `pwd`:/app -w /app scottkurt/gallery:cobertura /bin/sh -c "gocover-cobertura < cover.txt > coverage.xml"
    cd ../
done 