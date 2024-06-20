## This is a work in progress. !!!

This package will be a framework where you can build a code what you can deploy to 
    - AWS lambda
    - EC2 instance

Code once, deploy where you need
You can create a config file, which "handlers" you want to include in your build file. 
The command in the selective_builder will build the righ deployable executable for you.

This package contains a test docker image, using AWS localstack

Frontend to look into your S3, SNS and SQS queues

In the prebuilt_http function run it with go run main.go (not dot)

### Usage:

Create a config.yaml in `selective_builder` folder like this:
Here you can define which modules you want to group to lambda and wich to http package

```
port: 8080
lambda:
  packages:
    - name: example
      functions: 
        - route: /:TestHandler
    - name: example2
      functions: 
        - route: /add:TestHandler
http:
  packages:
    - name: example
      functions: 
        - route: /:TestHandler
    - name: example2
      functions: 
        - route: /add:TestHandler
        - route: /route2:TestHandler

```

Build it with the selective builder:
Usage:

```selective_builder <lambda|http>```

Exampe:

```
selective_builder http
selective_builder lambda

```

Each of them will create a source code specific to your environment:
- prebuildhttp
- prebuildlambda

You can deploy your lambda function via:
```
make build-deploy-lambda
```

Test your labdba withoud deploy:
```
make run-lambda
```

In your ```src/packages``` folder you need to create your packages, where you have to add handlers.(listed in your config.yaml)

The selective builder will verify if you missing any handler, incorectly specify your package name in your yaml

### Example:
```
// Package example is an example of a shared module handler
package example2

import (
	"context"
)

type Request struct {
	Name string `json:"name"`
}

// TestHandler is the unfied entry point of the module
func TestHandler(_ *context.Context, request *Request) (*Request, error) {
	return request, nil
}
```

### example with event dispatcher
```
// Package example is an example of a shared module handler
package example

import (
	"context"
	"fmt"

	dispather "attilaolbrich.co.uk/sqs_eventdispatcher"
)

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	ResponseName string `json:"respopnseName"`
}

// TestHandler is the unfied entry point of the module
func TestHandler(_ *context.Context, request *Request) (*Response, error) {

	fmt.Println(request)

	str, err := dispather.NewDispatcher().Send(*request)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)

	return &Response{
		ResponseName: "It is the response",
	}, nil
}

```

### The package main function looks like:
Here you keep the httpwrapper import, the lambda wrapper is for lambda built.


Whey you building with the ```selective builder``` then this file auto generated and will only contain the correct imports
and it will only list the proper handler definitions you privoded in the yaml file.

Only the used modules listed in your yaml will be copied to the prebuild folder(s)
```
// Package main is the main entry point
package main

import (
	"fmt"

	// connector "attilaolbrich.co.uk/lambdawrapper"

	"attilaolbrich.co.uk/example"
	"attilaolbrich.co.uk/example2"
	connector "attilaolbrich.co.uk/httpwrapper"
)

func main() {
	listener := connector.New()
	listener.Port(8080)
	err := listener.Start(
		connector.HandlerDef{Route: "/", Handler: example.TestHandler},
		connector.HandlerDef{Route: "/add", Handler: example2.TestHandler},
	)

	if err != nil {
		fmt.Println(err)
	}
}
```

### Test your SQS que visually:
```
http://localhost:8081/
```
where you can create SNS, SQS queues and look at the messages published, 
- Manage SNS subscriptions
- Manage S3
(all in your testing localstack)


### What is next

- Listener
- S3
- Full event driven architecture 
- and more...

