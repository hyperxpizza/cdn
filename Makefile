
current_dir=$(shell pwd)

build:
	go build -o bin/main ./cmd/cdn/main.go
gen_proto:
	protoc --go_out =./pkg/grpc --go_opt=paths=source_relative --go_out =./pkg/grpc --go-grpc_out=. --go-grpc_opt=paths=source_relative ./pkg/grpc/grpc.proto
psql:
	docker-compose exec postgres psql -d cdndb -U cdndbuser
docker_build:
	docker build -t cdn .
run:
	./bin/main --config=/home/hyperxpizza/dev/golang/cdn/config.json --grpc=true --rest=true
cert:
	openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout ./cert/nginx.key -out ./cert/nginx.cert  -subj '/CN=nginx'