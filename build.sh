#!/bin/bash

echo "Loading templates..."
cd load && go run template_loader.go
if [ $? -ne 0 ]; then { echo "Failed, aborting." ; exit 1; } fi
echo "Building..."
cd ../ && go build
if [ $? -ne 0 ]; then { echo "Failed, aborting." ; exit 1; } fi
echo "Finished!!!"
