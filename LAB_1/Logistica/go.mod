module Logistica

go 1.18

require (
	Caravanas v0.0.0-00010101000000-000000000000
	github.com/streadway/amqp v1.1.0
	google.golang.org/grpc v1.52.0
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.4.0 // indirect
	golang.org/x/sys v0.3.0 // indirect
	golang.org/x/text v0.5.0 // indirect
	google.golang.org/genproto v0.0.0-20221118155620-16455021b5e6 // indirect
)

replace Caravanas => ../Caravanas
