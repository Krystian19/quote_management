FROM golang:1.22.2-alpine3.19 AS builder

ENV GO111MODULE=on

RUN apk add make
RUN apk add git
RUN apk add gcc
RUN apk add libc-dev
RUN apk add pkgconfig

WORKDIR /go/src/app
COPY . /go/src/app

RUN make internal-bff-build

FROM alpine:3.19

WORKDIR /app
COPY --from=builder /go/src/app/build/internal-bff-app /app

EXPOSE 4000

ENTRYPOINT [ "./internal-bff-app" ]
