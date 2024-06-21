#!/bin/bash

cd ./selectivebuilder
go run . http
cd ../src/built/http
go run .
