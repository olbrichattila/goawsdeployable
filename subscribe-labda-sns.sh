
aws --endpoint-url=http://localhost:4566 sns create-topic --name my-topic

aws --endpoint-url=http://localhost:4566 sns subscribe \
    --topic-arn arn:aws:sns:us-east-1:000000000000:my-topic \
    --protocol lambda \
    --notification-endpoint arn:aws:lambda:us-east-1:000000000000:function:my-go-function-2
