module attilaolbrich.co.uk/example

go 1.22.4

replace attilaolbrich.co.uk/sqs_eventdispatcher => ../../eventdispatcher/sqs/

require attilaolbrich.co.uk/sqs_eventdispatcher v0.0.0-00010101000000-000000000000

require (
	github.com/aws/aws-sdk-go v1.54.4 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
)
