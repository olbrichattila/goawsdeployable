
## Work in Progress – Prototype

This is a work in progress and currently only a prototype.

An event created in SQS will be pulled by SNS. SNS supports subscriptions for Lambda functions and/or HTTP endpoints.  
The event will be processed by these applications in a unified way, and the developer doesn't need to know where the application is deployed.  
It is possible to create multiple Lambda packages and HTTP (EC2) packages in various combinations.

There is a prototype that writes to the database and executes a list of pre-defined instructions triggered by a single event.

This package serves as a framework for writing code that can be deployed to:
- AWS Lambda
- EC2 Instances

**Code once, deploy where needed.**  
You can create a `config.yaml` file specifying which "handlers" to include in your build.  
The `selective_builder` command will generate the correct deployable executable for you.

This package includes a test Docker image using AWS LocalStack and a frontend to inspect your S3, SNS, and SQS queues.

To run the prebuilt HTTP server, use:

```sh
go run main.go
```

### Usage

Create a `config.yaml` file in the `selective_builder` folder like this:  
Here you define which modules to group into Lambda or HTTP packages.

```yaml
port: 8080
lambda:
  packages:
    - name: example
      functions: 
        - snsroute: /:TestHandler
    - name: example2
      functions: 
        - route: /add:TestHandler
http:
  packages:
    - name: example
      functions: 
        - snsroute: /:TestHandler
    - name: example2
      functions: 
        - route: /add:TestHandler
        - route: /route2:TestHandler
```

Build using the selective builder:

```sh
selective_builder <lambda|http>
```

Example:

```sh
selective_builder http
selective_builder lambda
```

Deploy your Lambda function with:

```sh
make build-deploy-lambda
```

Test your Lambda locally without deploying:

```sh
make run-lambda
```

In your `src/packages` folder, create your packages and add handlers (listed in your `config.yaml`).  
The selective builder will verify if any handler is missing or incorrectly specified in the YAML.

### Example Handler (Basic)

```go
// Package example demonstrates creating a module usable in both AWS Lambda and HTTP with the same code
package example

import (
    "context"
    "sharedconfig"
    dispatcher "sqseventdispatcher"
)

type Request struct {
    Name string `json:"name"`
}

type Response struct {
    ResponseName string `json:"responseName"`
    ConfigType   string `json:"configType"`
}

// TestHandler is the unified entry point of the module
func TestHandler(_ *context.Context, config sharedconfig.SharedConfiger, request *Request) (*Response, error) {
    awsConfig := config.GetSQSConfig()
    err := dispatcher.NewDispatcher(awsConfig).Send(*request)
    if err != nil {
        return nil, err
    }

    return &Response{
        ResponseName: request.Name,
        ConfigType:   config.GetConfigType(),
    }, nil
}
```

### Example with Event Dispatcher

```go
package example

import (
    "context"
    "sharedconfig"
    dispatcher "sqseventdispatcher"
)

type Request struct {
    Name string `json:"name"`
}

type Response struct {
    DispatcherResult string `json:"dispatcherResult"`
    ResponseName     string `json:"responseName"`
    ConfigType       string `json:"configType"`
}

func TestHandler(_ *context.Context, config sharedconfig.SharedConfiger, request *Request) (*Response, error) {
    dispatcherResult, err := dispatcher.NewDispatcher().Send(*request)
    if err != nil {
        return nil, err
    }

    return &Response{
        ResponseName:     request.Name,
        DispatcherResult: dispatcherResult,
        ConfigType:       config.GetConfigType(),
    }, nil
}
```

### SNS Routing

If you define an `snsroute` (not just `route`) in your `config.yaml`, the Lambda and HTTP server will handle messages like this:

```go
package example

import (
    "context"
    "sharedconfig"
    "sqseventdispatcher"
)

type request struct {
    Name string `json:"name"`
}

type response struct {
    Request request `json:"request"`
}

func TestHandler(_ *context.Context, config sharedconfig.SharedConfiger, request *request) (*response, error) {
    dispatcher := sqseventdispatcher.NewDispatcher(config.GetSQSConfig())
    err := dispatcher.Send(*request)
    if err != nil {
        return nil, err
    }

    return &response{
        Request: *request,
    }, nil
}
```

### `main.go` Entry Point

When building with the `selective_builder`, this file is auto-generated with the correct imports and only includes the handler definitions specified in the YAML file.

Only the required modules listed in your YAML will be copied to the prebuild folder(s).

```go
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

### Visual Testing for SQS Queue

Visit:

```
http://localhost:8081/
```

You can:
- Create SNS and SQS queues
- View published messages
- Manage SNS subscriptions
- Manage S3 (all via LocalStack for testing)

### What's Next

- Listener
- S3 integration
- Full event-driven architecture
- ...and more!
