#!/bin/bash 
cd ../
echo "building packages"
sh Dev-build-packages.sh
echo "generating coverage"
sh Dev-generate-coverage.sh