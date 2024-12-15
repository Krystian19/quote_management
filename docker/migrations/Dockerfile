FROM golang:1.22.2-alpine3.19 

ENV GO111MODULE=on
ENV GOOSE_DRIVER="postgres"

# Install env dependencies
RUN apk add make
RUN apk add git
RUN apk add gcc
RUN apk add libc-dev

WORKDIR /go/src/app
COPY . /go/src/app

RUN go install -mod=vendor ./vendor/github.com/pressly/goose/v3/cmd/goose

