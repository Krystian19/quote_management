build: \
	external-bff-build \
	internal-bff-build \

test:
	go test ./internal/... -timeout 20s

test-no-cache:
	go test ./internal/... -count=1 -timeout 20s

gqlgen:
	(go run github.com/99designs/gqlgen@v0.17.49)
	(go run github.com/Khan/genqlient@v0.7.0)

proto-gen:
	protoc  \
	--go_out=. \
	--go-grpc_opt=require_unimplemented_servers=false \
	--go_opt=paths=source_relative \
	--go-grpc_out=. \
	--go-grpc_opt=paths=source_relative ./proto/*.proto \

external-bff-build:
	go build \
		-ldflags "-s -w" \
		-gcflags "all=-trimpath=${PWD}" \
		-tags netgo \
		-o build/external-bff-app ./cmd/external-bff/*.go

internal-bff-build:
	go build \
		-ldflags "-s -w" \
		-gcflags "all=-trimpath=${PWD}" \
		-tags netgo \
		-o build/internal-bff-app ./cmd/internal-bff/*.go
