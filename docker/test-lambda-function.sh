aws --endpoint-url=http://localhost:4566 lambda invoke \
    --function-name my-go-function-2 \
    --payload '{"path": "/", "name": "World", "items": ["item1", "item2"]}' \
    response.json

cat response.json
echo ""
