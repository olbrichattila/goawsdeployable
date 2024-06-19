module attilaolbrich.co.uk/deployable

go 1.22.4

replace attilaolbrich.co.uk/httpwrapper => ../deployment_wrapper/http

replace attilaolbrich.co.uk/lambdawrapper => ../deployment_wrapper/lambda

replace attilaolbrich.co.uk/sqs_event_dispatcher => ../event_dispatcher/sqs/

replace attilaolbrich.co.uk/example => ./example

replace attilaolbrich.co.uk/example2 => ./example2

replace attilaolbrich.co.uk/handler => ../handler


require (
	attilaolbrich.co.uk/example v0.0.0-00010101000000-000000000000
	attilaolbrich.co.uk/example2 v0.0.0-00010101000000-000000000000
	attilaolbrich.co.uk/httpwrapper v0.0.0-00010101000000-000000000000
)

require (
	github.com/aws/aws-sdk-go v1.54.4 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	attilaolbrich.co.uk/handler v0.0.0-00010101000000-000000000000 // indirect
	attilaolbrich.co.uk/sqs_event_dispatcher v0.0.0-00010101000000-000000000000 // indirect
)
