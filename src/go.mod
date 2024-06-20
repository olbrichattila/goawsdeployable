module attilaolbrich.co.uk/deployment

go 1.22.4

replace attilaolbrich.co.uk/routebuilder => ./routebuilder

replace attilaolbrich.co.uk/deploymentwrapper/http => ./deploymentwrapper/http

replace attilaolbrich.co.uk/deploymentwrapper/lambda => ./deploymentwrapper/lambda

replace attilaolbrich.co.uk/example => ./packages/example

replace attilaolbrich.co.uk/example2 => ./packages/example2

replace attilaolbrich.co.uk/eventdispatcher/sqs => ./eventdispatcher/sqs

replace attilaolbrich.co.uk/handler => ./handler

require attilaolbrich.co.uk/routebuilder v0.0.0-00010101000000-000000000000

require (
	attilaolbrich.co.uk/deploymentwrapper/http v0.0.0-00010101000000-000000000000 // indirect
	attilaolbrich.co.uk/eventdispatcher/sqs v0.0.0-00010101000000-000000000000 // indirect
	attilaolbrich.co.uk/example v0.0.0-00010101000000-000000000000 // indirect
	attilaolbrich.co.uk/example2 v0.0.0-00010101000000-000000000000 // indirect
	attilaolbrich.co.uk/handler v0.0.0-00010101000000-000000000000 // indirect
	github.com/aws/aws-sdk-go v1.54.6 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
)
