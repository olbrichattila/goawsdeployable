aws --endpoint-url=http://localhost:4566 lambda invoke \
    --function-name my-go-function-2 \
    --payload '{"path": "/", "name": "hello from lambda"}' \
    response.json

cat response.json
echo ""
