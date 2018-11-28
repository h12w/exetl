generate:
		protoc -I proto/ proto/storage.proto --go_out=plugins=grpc:proto

lint:
		golint ./...

test:
		go test -timeout=10s ./...