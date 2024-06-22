#!/bin/bash

IMAGE_NAME=my-go-app

cd ./http-logger
go build -o main main.go
cp ./main ../docker/http/main
cd ../docker/http

docker stop $IMAGE_NAME || true
docker rm -f $IMAGE_NAME || true
docker rmi -f $IMAGE_NAME || true
docker build -t $IMAGE_NAME .
docker run -d --name my-go-app --network docker_my-dev -p 8080:8080 $IMAGE_NAME
