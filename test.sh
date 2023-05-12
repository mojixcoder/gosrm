#!/bin/bash

# Default values for testing.
osrm_address="http://127.0.0.1:5000"
coverprofile=false

while getopts "a:p" flags; do
    case "${flags}" in
        a) 
            osrm_address=${OPTARG}
            ;;
        p) 
            coverprofile=true
            ;;
    esac
done

export OSRM_ADDRESS=$osrm_address
echo "OSRM_ADDRESS: $OSRM_ADDRESS"

if [ $coverprofile = true ]; then
    go test -v -coverprofile=coverage.out $(go list ./... | grep -v /examples)
else
    go test -v -cover $(go list ./... | grep -v /examples)
fi
