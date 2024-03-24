#!/bin/bash
services=`ls  | grep service`
for service in $services 
do 
    cd $service 
    go test -coverprofile=coverage.txt
    gocover-cobertura < coverage.txt > coverage.xml
    cd ../
done 