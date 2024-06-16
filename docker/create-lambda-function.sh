
#!/bin/bash
aws --endpoint-url=http://localhost:4566 lambda create-function \
    --function-name my-go-function-2 \
    --runtime go1.x \
    --role arn:aws:iam::000000000000:role/lambda-role \
    --handler main \
    --zip-file fileb://../built/lambda/function.zip
