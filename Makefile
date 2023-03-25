server:
	go run cmd/main.go

test:
	go test -v -cover -short ./...

proto:
	protoc --proto_path=proto --go_out=proto/pb --go_opt=paths=source_relative \
		--go-grpc_out=proto/pb --go-grpc_opt=paths=source_relative \
		proto/*.proto

.PHONY: server test proto