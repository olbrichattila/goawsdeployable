cd docker
docker-compose restart
./create-lambda-function.sh 
sleep 2
./test-lambda-function.sh 
