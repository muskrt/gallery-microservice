#!/bin/bash
services=`ls  | grep '_service'`
for service in $services 
do 
    cd $service 
    go test -coverprofile=coverage.txt
    /home/ec2-user/go/bin/gocover-cobertura < coverage.txt > coverage.xml
    cd ../
done 