module olbrichattila.co.uk/test

go 1.18

replace olbrichattila.co.uk/httpwrapper => ../deployment_wrapper/http

replace olbrichattila.co.uk/lambdawrapper => ../deployment_wrapper/lambda

replace olbrichattila.co.uk/example => ./example

replace olbrichattila.co.uk/example2 => ./example2

require (
	olbrichattila.co.uk/example v0.0.0-00010101000000-000000000000
	olbrichattila.co.uk/example2 v0.0.0-00010101000000-000000000000
	olbrichattila.co.uk/httpwrapper v0.0.0-00010101000000-000000000000
)
