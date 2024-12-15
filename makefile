build: \
	external-bff-build \

test:
	go test ./internal/... -timeout 20s

test-no-cache:
	go test ./internal/... -count=1 -timeout 20s

external-bff-build:
	go build \
		-ldflags "-s -w" \
		-gcflags "all=-trimpath=${PWD}" \
		-tags netgo \
		-o build/external-bff-app ./cmd/external-bff/*.go
