cd docker
docker-compose restart
./create-lambda-function.sh 
sleep 2
aws --endpoint-url=http://localhost:4566 sqs create-queue --queue-name test
./test-lambda-function.sh 
