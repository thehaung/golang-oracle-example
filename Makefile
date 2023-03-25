server:
	go run cmd/main.go

proto:
	protoc --proto_path=proto --go_out=proto/pb --go_opt=paths=source_relative \
		--go-grpc_out=proto/pb --go-grpc_opt=paths=source_relative \
		proto/*.proto

.PHONY: server proto