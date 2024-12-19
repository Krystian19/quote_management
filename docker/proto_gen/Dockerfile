FROM golang:1.22.2-alpine3.19

ENV GO111MODULE=on

RUN apk add make
RUN apk add git
RUN apk add protobuf-dev

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.2
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.4

WORKDIR /go/src/app
