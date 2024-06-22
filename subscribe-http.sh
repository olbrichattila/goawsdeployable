#!/bin/bash

aws --endpoint-url=http://localhost:4566 sns subscribe \
    --topic-arn arn:aws:sns:us-east-1:000000000000:my-topic \
    --protocol http \
    --notification-endpoint http://my-go-app:8080