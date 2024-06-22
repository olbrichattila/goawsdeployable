aws-list-functions:
	aws --endpoint-url=http://localhost:4566 lambda list-functions --query 'Functions[*].FunctionArn' --output text
setup: build-deploy-lambda build-http
	./subscribe-labda-sns.sh
	./subscribe-http.sh
build-deploy-lambda: build-lambda deploy-lambda
build-lambda:
	./build-lambda.sh
deploy-lambda:
	./redeploy-lambda.sh
run-lambda:
	 ./docker/test-lambda-function.sh
build-http:
	./build-http.sh
run-http:
	./run-http.sh
lint:
	gocritic check ./...
	revive ./...
	golint ./...
	goconst ./...
	golangci-lint run
	go vet ./...
	staticcheck ./...
test:
	go test ./...
