package lambda_listener

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func New() Listener {
	return &listen{}
}

type Listener interface {
	Start(handlers ...HandlerDef) error
	Port(int)
}

type listen struct {
}

type HandlerDef struct {
	Route   string
	Handler HandlerFunc
}
type HandlerFunc = func(ctx context.Context, payload string) (string, error)
type lambdaHandlerFunc = func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

func (l *listen) Start(handlers ...HandlerDef) error {
	lambda.Start(l.middleware(handlers...))

	return nil
}

func (l *listen) Port(_ int) {
	// this function is not requred for Lambda
}

func (l *listen) middleware(handlers ...HandlerDef) lambdaHandlerFunc {
	return func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		var res string
		var err error
		path := request.Path
		if path == "" {
			path = "/"
		}

		// @TODO this can be expanded with many paraeters from APIGatewayProxyRequest, like HTTPMethod, headers, QueryStringParameters, PathParameters
		for _, fcHandler := range handlers {
			if path == fcHandler.Route {
				res, err = fcHandler.Handler(ctx, request.Body)
				if err != nil {
					return events.APIGatewayProxyResponse{
						StatusCode: 500,
						Body:       err.Error(),
					}, nil
				}

				return events.APIGatewayProxyResponse{
					StatusCode: 200,
					Body:       res,
				}, nil
			}
		}
		return events.APIGatewayProxyResponse{
			StatusCode: 404,
			Body:       "Route not found",
		}, nil
	}
}
