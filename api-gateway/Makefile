proto:
	protoc pkg/**/pb/*.proto --go_out=. --go-grpc_out=.
	@echo "Protoc Compile Successfully"

server:
	go run cmd/main.go