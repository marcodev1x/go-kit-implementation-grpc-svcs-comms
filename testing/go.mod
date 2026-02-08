module github.com/testing

go 1.24.4

require (
	github.com/go-kit/kit v0.13.0
	github.com/go-kit/log v0.2.0
	github.com/marco-kit/kit-home-service v0.0.0
	github.com/oklog/run v1.2.0
	google.golang.org/grpc v1.78.0
)

require (
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	golang.org/x/net v0.49.0 // indirect
	golang.org/x/sys v0.40.0 // indirect
	golang.org/x/text v0.33.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260128011058-8636f8732409 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)

replace github.com/marco-kit/kit-home-service => ../kit-home-service
