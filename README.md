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
// Package example2 is just an example how to create a module accross AWS lambda or HTTP with the same code
package example2

import (
	"context"
	"fmt"
	"sharedconfig"
)

// The Request automatically marshalled here
type Request struct {
	Name string `json:"name"`
}

// The response, which is unmarshalled automatically and returned in http or lambda response
type Response struct {
	ConfigType  string `json:"configType"`
	RequestName string `json:"requestName"`
}

// TestHandler is the unfied entry point of the module
func TestHandler(_ *context.Context, config sharedconfig.SharedConfiger, request *Request) (*Response, error) {
	fmt.Println()

	return &Response{
		ConfigType:  config.GetConfigType(),
		RequestName: request.Name,
	}, nil
}
```

### example with event dispatcher
```
// Package example is just an example how to create a module accross AWS lambda or HTTP with the same code
package example

import (
	"context"
	"sharedconfig"

	dispather "sqseventdispatcher"
)

// Request data automaticall marhalled here
type Request struct {
	Name string `json:"name"`
}

// Response what we want to be returned as HTTP or Lambda
type Response struct {
	DispactherResult string `json:"dispatherResult"`
	ResponseName     string `json:"respopnseName"`
	ConfigType       string `json:"configType"`
}

// TestHandler is the unfied entry point of the module
func TestHandler(_ *context.Context, config sharedconfig.SharedConfiger, request *Request) (*Response, error) {
	dispatcherResult, err := dispather.NewDispatcher().Send(*request)
	if err != nil {
		return nil, err
	}

	return &Response{
		ResponseName:     request.Name,
		DispactherResult: dispatcherResult,
		ConfigType:       config.GetConfigType(),
	}, nil
}
```

### The package main function looks like:
Here you keep the httpwrapper import, the lambda wrapper is for lambda built.


Whey you building with the ```selective builder``` then this file auto generated and will only contain the correct imports
and it will only list the proper handler definitions you privoded in the yaml file.

Only the used modules listed in your yaml will be copied to the prebuild folder(s)
```
// Package main is the entry point
package main

import (
	"fmt"
	"deploymentwrapper"
	"example"
	"example2"
	connector "httplistener"
	config "httpconfig"
)

func main() {
	listener := connector.New()
	listener.Config(config.New())
	listener.Port(8080)
	err := listener.Start(
		deploymentwrapper.HandlerDef{Route: "/", Handler: example.TestHandler},
		deploymentwrapper.HandlerDef{Route: "/add", Handler: example2.TestHandler},
		deploymentwrapper.HandlerDef{Route: "/route2", Handler: example2.TestHandler},
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

