#!/bin/bash

# NOTE: Currently there is a LIBC6 version issue, as localstack is not up-to date
# therefore the app compiled with  CGO_ENABLED=0. this can be overcome by deploying to lambda with docker.

dir="./built/lambda/"
file="./built/lambda/function.zip"

# Check if directory exists
if [ ! -d "$dir" ]; then
    # Directory does not exist, create it
    mkdir -p "$dir"
    echo "Directory $dir created."
else
    echo "Directory $dir already exists."
fi

if [ -f "$file" ]; then
    # Delete the file
    rm "$file"
    echo "File '$file' deleted."
else
    echo "File '$file' does not exist."
fi

cd src/packages
# Initialize the module if not already done
go mod init olbrichattila.co.uk || true

go mod tidy

if [ -f ./main ]; then
    # Delete the file
    rm ./main
    echo "File main deleted."
else
    echo "File main does not exist."
fi

# Build the Go executable for Linux (required by AWS Lambda)
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main main.go

chmod +x main
# Package the executable into a zip file
zip ./../../built/lambda/function.zip main
