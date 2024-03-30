#!/bin/bash
echo "Running build script "
echo $PWD
services=`ls  | grep '_service'`
for service in $services 
do 
cd $service 
docker run --rm -v `pwd`:/app -w /app golang:1.22-alpine  /bin/sh -c "go build . "
cd ../
done