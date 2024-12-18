FROM golang:1.22.2-alpine3.19 AS builder

ENV GO111MODULE=on

RUN apk add make
RUN apk add git
RUN apk add gcc
RUN apk add libc-dev
RUN apk add pkgconfig

WORKDIR /go/src/app
COPY . /go/src/app

RUN make external-bff-build

FROM alpine:3.19

WORKDIR /app
COPY --from=builder /go/src/app/build/external-bff-app /app

EXPOSE 4000

ENTRYPOINT [ "./external-bff-app" ]
