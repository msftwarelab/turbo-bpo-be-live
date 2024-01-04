#!/bin/bash

echo -e "Creating binary.."

env GOOS=linux GOARCH=amd64 GO111MODULE=on /usr/local/go/bin/go build -a -installsuffix cgo -mod vendor -o ./main-jobs-ping-staging
if [[ $? -eq 0 ]]
then
  echo -e "Binary was created successfully"
fi
  