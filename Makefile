build:
	go build -o bin/main ./cmd/cdn/main.go
gen_proto:
	protoc --go_out =./pkg/grpc --go_opt=paths=source_relative -go_out =./pkg/grpc --go-grpc_out=. --go-grpc_opt=paths=source_relative grpc.proto
