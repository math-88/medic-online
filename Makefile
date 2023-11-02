.PHONY: all server test proto build run compile sql compose up

proto:
	buf generate --template proto/v1/buf.gen.yaml proto/v1 -o proto/v1
# protoc --go_out=. --go_opt=paths=source_relative \
# --go-grpc_out=. --go-grpc_opt=paths=source_relative \
# helloworld/helloworld.proto

protoc -I. --proto_path=proto/v1/ --graphql_out=proto/v1/gen/go proto/v1/auth.proto

sql:
	sqlc generate

build:
	go build -o bin/medical cmd/main.go

mod:
	go mod tidy

run:
	go run cmd/main.go

compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o bin/main-linux-arm cmd/main.go
	GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 cmd/main.go
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 cmd/main.go

test:
	go test -cover ./...

all: proto sql mod test run
server: run
compose: docker compose up -d
up: proto sql test compose