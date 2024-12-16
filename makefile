build: \
	external-bff-build \

test:
	go test ./internal/... -timeout 20s

test-no-cache:
	go test ./internal/... -count=1 -timeout 20s

gqlgen:
	(go run github.com/99designs/gqlgen@v0.17.49)

external-bff-build:
	go build \
		-ldflags "-s -w" \
		-gcflags "all=-trimpath=${PWD}" \
		-tags netgo \
		-o build/external-bff-app ./cmd/external-bff/*.go
